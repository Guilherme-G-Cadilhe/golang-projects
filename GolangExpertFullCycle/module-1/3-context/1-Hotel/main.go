package main

import (
	"context"
	"fmt"
	"time"
)

/* Cria e gera um Context que é uma forma de controlar a vida útil de métodos
que vão ser executados em background e que podera ser cancelado.
Ex: Timeout, Cancel, Deadline, Value
Muito utilizado para controlar o tempo de resposta de uma requisição HTTP
*/

func main() {

	// Contexto em Branco ( inicializar )
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {

	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking canceled. Timeout Reached")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel booked")
		return

	}

}
