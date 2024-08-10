package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type APIRETURN struct {
	USDBRL USDBRL `json:"USDBRL"`
}
type USDBRL struct {
	Code        string `json:"code"`
	Codein      string `json:"codein"`
	Name        string `json:"name"`
	High        string `json:"high"`
	Low         string `json:"low"`
	VarBid      string `json:"varBid"`
	PctChange   string `json:"pctChange"`
	Bid         string `json:"bid"`
	Ask         string `json:"ask"`
	Timestamp   string `json:"timestamp"`
	Create_date string `json:"create_date"`
}

func main() {
	// os.Remove("./cotacao.db")
	db, err := sql.Open("sqlite3", "./cotacao.db")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer db.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /cotacao", func(w http.ResponseWriter, r *http.Request) {
		ConsultaCotacao(db, w, r)
	})
	http.ListenAndServe(":8080", mux)
}

func ConsultaCotacao(db *sql.DB, response http.ResponseWriter, request *http.Request) {
	// Padrão para o Exercicio é 200ms
	ctxApi, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	result, err := FetchCotacaoUSDBRL(ctxApi)
	if err != nil {
		if handleContextError(ctxApi, response, err, "API Cambio demorou demais") {
			return
		}
		handleHttpError(response, http.StatusInternalServerError, "Erro buscando cotação", err)
		return
	}
	if result == nil || result.USDBRL == (USDBRL{}) {
		handleHttpError(response, http.StatusInternalServerError, "Dados de câmbio não encontrados", nil)
		return
	}

	// Padrão para o Exercicio é 10ms
	ctxDb, cancelDb := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelDb()

	err = saveCotacaoOnBD(ctxDb, db, result.USDBRL)
	if err != nil {
		if handleContextError(ctxDb, response, err, "Banco de dados demorou demais") {
			return
		}
		handleHttpError(response, http.StatusInternalServerError, "Erro ao salvar cotação", err)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode(map[string]string{"cotacao": result.USDBRL.Bid})
	if err != nil {
		handleHttpError(response, http.StatusInternalServerError, "Erro ao gerar resposta JSON:", err)
		return
	}
}

func FetchCotacaoUSDBRL(ctx context.Context) (*APIRETURN, error) {

	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	requestConfig, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(requestConfig)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, error := io.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}
	var data APIRETURN
	error = json.Unmarshal(body, &data)
	if error != nil {
		return nil, error
	}
	return &data, nil
}

func saveCotacaoOnBD(ctx context.Context, db *sql.DB, cotacao USDBRL) error {
	_, err := db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS cotacoes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			inserted_at DATETIME NOT NULL,
			code TEXT,
			codein TEXT,
			name TEXT,
			high TEXT,
			low TEXT,
			var_bid TEXT,
			pct_change TEXT,
			bid TEXT,
			ask TEXT,
			timestamp TEXT,
			create_date TEXT
		)
	`)
	if err != nil {
		return err
	}

	stmt, err := db.PrepareContext(ctx, `
		INSERT INTO cotacoes (
			inserted_at, code, codein, name, high, low, var_bid, pct_change, bid, ask, timestamp, create_date
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		time.Now(), cotacao.Code, cotacao.Codein, cotacao.Name, cotacao.High, cotacao.Low,
		cotacao.VarBid, cotacao.PctChange, cotacao.Bid, cotacao.Ask, cotacao.Timestamp, cotacao.Create_date)
	return err
}

func handleContextError(ctx context.Context, response http.ResponseWriter, err error, customMessage string) bool {
	if ctx.Err() == context.DeadlineExceeded {
		handleHttpError(response, http.StatusGatewayTimeout, customMessage, err)
		return true
	}
	return false
}
func handleHttpError(response http.ResponseWriter, statusCode int, customMessage string, err error) {
	fmt.Println("[Erro]:", customMessage, err)

	errorMessage := customMessage
	if customMessage == "" && err != nil {
		errorMessage = err.Error()
	}

	response.WriteHeader(statusCode)
	json.NewEncoder(response).Encode(map[string]string{"error": errorMessage})
}
