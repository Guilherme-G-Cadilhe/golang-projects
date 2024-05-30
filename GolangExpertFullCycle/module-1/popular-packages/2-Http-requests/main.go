package main

import (
	"io"
	"net/http"
)

func main() {

	// Criação de uma requisição HTTP
	request, err := http.Get("https://www.google.com.br")
	if err != nil {
		panic(err)
	}
	// Encerra a requisição usando defer, que é executado após o final do bloco de código, por último.
	defer request.Body.Close()

	// Leitura da resposta HTTP
	result, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	// Encerra a requisição, da forma antiga e manual
	// request.Body.Close()

	// Transforma a resposta de bytes em string
	println(string(result))

}
