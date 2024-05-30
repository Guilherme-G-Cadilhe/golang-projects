package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// https://mholt.github.io/json-to-go/
// https://transform.tools/json-to-go
// go run main.go https://viacep.com.br/ws/01001000/json ou N argumentos no terminal
// go run main.go 01001000 79304-411 ou N argumentos no terminal
// go build -o cep.exe main.go
// ./cep.exe 79304-411

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Uso: %s <CEP>...\n", os.Args[0])
		os.Exit(1)
	}

	for index, cep := range os.Args[1:] {
		fmt.Println("Consultando CEP:", cep)
		url := "https://viacep.com.br/ws/" + cep + "/json/"
		fmt.Println("URL:", url)
		request, err := http.Get(url)
		if err != nil {
			// Fprint joga os dados para a saida padrao (Stdout = saida padrao , terminal)
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v", err)
			// panic(err)
			continue
		}
		defer request.Body.Close()

		// readAll le o corpo da requisição até o final e retorna um slice de bytes (ByteArray)
		result, err := io.ReadAll(request.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v", err)
			// panic(err)
			continue
		}
		fmt.Println(string(result))

		var data ViaCEP
		err = json.Unmarshal(result, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao decodificar JSON: %v", err)
			// panic(err)
			continue
		}
		fmt.Println("Data:", data, "Localidade:", data.Localidade)

		if data.Logradouro == "" {
			data.Logradouro = fmt.Sprintf("Desconhecido%v", index)
		}

		fileName := data.Logradouro + ".txt"
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v", err)
			// panic(err)
			continue
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("%v\nLocalidade: %v\n", data, data.Localidade))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v", err)
			panic(err)
		}

	}
}
