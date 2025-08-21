package main

import (
	"fmt"
	"time"
)

func main() {
	// Canal com buffer para 3 mensagens
	ch := make(chan int, 3)

	// O emissor envia 3 mensagens rapidamente sem bloquear
	fmt.Println("Enviando 1...")
	ch <- 1
	fmt.Println("Enviando 2...")
	ch <- 2
	fmt.Println("Enviando 3...")
	ch <- 3
	fmt.Println("Todos os 3 valores enviados para o buffer.")

	// O receptor consome os valores lentamente
	time.Sleep(2 * time.Second)
	fmt.Println("Recebido:", <-ch)
	time.Sleep(1 * time.Second)
	fmt.Println("Recebido:", <-ch)
	time.Sleep(1 * time.Second)
	fmt.Println("Recebido:", <-ch)
}
