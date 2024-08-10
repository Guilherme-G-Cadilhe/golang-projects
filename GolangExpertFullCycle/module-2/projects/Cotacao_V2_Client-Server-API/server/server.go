package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"
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
	result, err := FetchCotacaoUSDBRL()
	if err != nil {
		handleHttpError(response, http.StatusInternalServerError, "Erro buscando cotação", err)
		return
	}
	if result == nil || result.USDBRL == (USDBRL{}) {
		handleHttpError(response, http.StatusInternalServerError, "Dados de câmbio não encontrados", nil)
		return
	}

	ctxDb, cancelDb := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancelDb()

	err = saveCotacaoOnBD(ctxDb, db, result.USDBRL)
	if err != nil {
		handleHttpError(response, http.StatusInternalServerError, "Erro ao salvar cotação", err)
		return
	}

	// printAllCotacoes(db)

	// Exibir os dados salvos na tabela
	records, err := fetchAllCotacoes(db)
	if err != nil {
		handleHttpError(response, http.StatusInternalServerError, "Erro ao consultar cotação", err)
		return
	}
	// printRecordsDelimitador(records)
	printRecordsLinhaNova(records)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode(map[string]string{"cotacao": result.USDBRL.Bid})
	if err != nil {
		handleHttpError(response, http.StatusInternalServerError, "Erro ao gerar resposta JSON:", err)
		return
	}
}

func handleHttpError(response http.ResponseWriter, statusCode int, message string, err error) {
	fmt.Println("[Erro]:", message, err)

	errorMessage := message
	if err != nil && strings.Contains(err.Error(), "context deadline exceeded") {
		errorMessage = "Context deadline exceeded"
	} else if err != nil {
		errorMessage = fmt.Sprintf("%s %s", errorMessage, err.Error())
	}

	response.WriteHeader(statusCode)
	json.NewEncoder(response).Encode(map[string]string{"error": errorMessage})
}

func FetchCotacaoUSDBRL() (*APIRETURN, error) {
	// Alterar os 300ms pra o primeiro hit não travar por lag
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

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

	// _, err = stmt.ExecContext(ctx,
	// 	time.Now(), cotacao.Code, cotacao.Codein, cotacao.Name, cotacao.High, cotacao.Low,
	// 	cotacao.VarBid, cotacao.PctChange, cotacao.Bid, cotacao.Ask, cotacao.Timestamp, cotacao.Create_date)
	values := getUSDBRLValues(cotacao)

	// Adiciona a data e hora atual
	newValues := append([]interface{}{time.Now()}, values...)

	_, err = stmt.ExecContext(ctx, newValues...)
	return err
}

func getUSDBRLValues(cotacao USDBRL) []interface{} {
	val := reflect.ValueOf(cotacao)
	values := make([]interface{}, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		values[i] = val.Field(i).Interface()
	}
	return values
}

// Função para buscar todos os registros na tabela cotacoes
func fetchAllCotacoes(db *sql.DB) ([]map[string]interface{}, error) {
	rows, err := db.Query("SELECT * FROM cotacoes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	results := []map[string]interface{}{}
	for rows.Next() {
		columnValues := make([]interface{}, len(columns))
		columnPointers := make([]interface{}, len(columns))
		for i := range columnValues {
			columnPointers[i] = &columnValues[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		rowResult := map[string]interface{}{}
		for i, colName := range columns {
			val := columnValues[i]
			rowResult[colName] = val
		}
		results = append(results, rowResult)
	}

	return results, nil
}

func printRecordsLinhaNova(records []map[string]interface{}) {
	for i, record := range records {
		fmt.Printf("Registro %d:\n", i+1)
		for key, value := range record {
			fmt.Printf("\t%s: %v\n", key, value)
		}
		fmt.Println() // Linha em branco entre registros
	}
}

func printRecordsDelimitador(records []map[string]interface{}) {
	for i, record := range records {
		fmt.Printf("Registro %d: ", i+1)
		parts := make([]string, 0, len(record))
		for key, value := range record {
			parts = append(parts, fmt.Sprintf("%s: %v", key, value))
		}
		fmt.Println(strings.Join(parts, " | "))
	}
}

// func printAllCotacoes(db *sql.DB) error {
// 	rows, err := db.Query("SELECT * FROM cotacoes")
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()

// 	columns, err := rows.Columns()
// 	if err != nil {
// 		return err
// 	}

// 	for rows.Next() {
// 		columnPointers := make([]interface{}, len(columns))
// 		for i := range columnPointers {
// 			columnPointers[i] = new(interface{})
// 		}

// 		if err := rows.Scan(columnPointers...); err != nil {
// 			return err
// 		}

// 		for i, colName := range columns {
// 			fmt.Printf("%s: %v  ", colName, *(columnPointers[i].(*interface{})))
// 		}
// 		fmt.Println()
// 	}
// 	return nil
// }
