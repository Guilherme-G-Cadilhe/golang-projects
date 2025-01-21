package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Structs para armazenar as respostas das APIs
type ViaCepResponse struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

type BrasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func fetchViaCep(cep string, ch chan<- ViaCepResponse, errCh chan<- error) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		errCh <- err
		return
	}
	defer resp.Body.Close()

	var data ViaCepResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		errCh <- err
		return
	}
	ch <- data
}

func fetchBrasilAPI(cep string, ch chan<- BrasilAPIResponse, errCh chan<- error) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	resp, err := http.Get(url)
	if err != nil {
		errCh <- err
		return
	}
	defer resp.Body.Close()

	var data BrasilAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		errCh <- err
		return
	}
	ch <- data
}

func main() {
	cep := "21545310"
	viaCepCh := make(chan ViaCepResponse)
	brasilAPICh := make(chan BrasilAPIResponse)
	errCh := make(chan error)

	// Inicia as requisições simultaneamente com GoRoutine
	go fetchViaCep(cep, viaCepCh, errCh)
	go fetchBrasilAPI(cep, brasilAPICh, errCh)

	select {
	case res := <-viaCepCh:
		fmt.Printf("Resposta da ViaCep: %+v\n", res)
	case res := <-brasilAPICh:
		fmt.Printf("Resposta da BrasilAPI: %+v\n", res)
	case err := <-errCh:
		fmt.Printf("Erro: %v\n", err)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: nenhuma resposta em 1 segundo")
	}
}
