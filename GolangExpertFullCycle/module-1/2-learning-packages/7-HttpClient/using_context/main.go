package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	/* Cria e gera um Context que é uma forma de controlar a vida útil de métodos
	que vão ser executados em background e que podera ser cancelado.
	Ex: Timeout, Cancel, Deadline, Value
	Muito utilizado para controlar o tempo de resposta de uma requisição HTTP
	*/

	// Cria um contexto vazio
	ctx := context.Background()

	// WithTimeout = cria um contexto com um timeout que cancela a requisição
	// Ele executa no sentido 1, 2 ,3 ,4 ,5
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond) // time.Second

	// WithDeadline = cria um contexto com um deadline que cancela a requisição
	// Ele executa no sentido 5, 4 ,3 ,2 ,1
	// ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))

	// WithValue = cria um contexto com um valor que pode ser recuperado quando o contexto for cancelado
	// ctx := context.WithValue(ctx, "key", "value")

	// WithCancel = cria um contexto que pode ser cancelado em qualquer momento
	// ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", "https://google.com", nil)
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
