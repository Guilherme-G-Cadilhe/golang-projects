package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type COTACAOAPI struct {
	Cotacao string `json:"cotacao"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Requisição demorou demais")
			panic(err)
		} else {
			fmt.Println("Erro ao fazer requisição:", err)
			panic(err)
		}

	}
	defer res.Body.Close()
	result, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var data COTACAOAPI
	err = json.Unmarshal(result, &data)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("cotacao.txt", []byte(fmt.Sprintf("Dólar: %s", data.Cotacao)), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Cotação salva com sucesso!")
	return
	// file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// _, err = file.WriteString(fmt.Sprintf("Dólar: %s\n", data.Cotacao))
	// if err != nil {
	// 	panic(err)

	// }

}
