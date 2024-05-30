package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//  curl localhost:8080/?cep=79304-411 -v

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
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		response.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := request.URL.Query().Get("cep")
	if cepParam == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	cep, error := BuscaCep(cepParam)
	if error != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	// Retorno resumido convertendo diretamente em JSON dentro do response (Writer = escritor)
	json.NewEncoder(response).Encode(cep)

	// // Retorno verboso fazendo as conversões por etapas e finalizando com o retorno tratado como JSON
	// jsonCEP, error := json.Marshal(cep)
	// if error != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// response.Write(jsonCEP)

}

func BuscaCep(cep string) (*ViaCEP, error) {
	fmt.Println("Consultando CEP:", cep)
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	fmt.Println("URL:", url)
	request, error := http.Get(url)
	if error != nil {
		return nil, error
	}
	defer request.Body.Close()
	body, error := io.ReadAll(request.Body)
	if error != nil {
		return nil, error
	}
	var data ViaCEP
	error = json.Unmarshal(body, &data)
	if error != nil {
		return nil, error
	}
	return &data, nil
}
