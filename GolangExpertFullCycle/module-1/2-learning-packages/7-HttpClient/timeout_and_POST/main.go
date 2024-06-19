package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	client := http.Client{
		// Timeout: time.Second,
	}

	responseGet, err := client.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err)
	}
	defer responseGet.Body.Close()

	body, err := io.ReadAll(responseGet.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	jsonVar := bytes.NewBuffer([]byte(`{"foo": "bar"}`))
	responsePost, err := client.Post("https://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer responsePost.Body.Close()

	// Pega os dados, escolhe para onde vai jogar e de onde é copiado, também seta um limite de dados se quiser
	io.CopyBuffer(os.Stdout, responsePost.Body, nil)

	// body, err = io.ReadAll(responsePost.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(body))

}
