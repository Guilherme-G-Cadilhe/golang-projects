package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("Goroutine: Preparando para enviar...")
		time.Sleep(2 * time.Second)
		ch <- "Olá, Main!" // Bloqueia aqui até a main receber
		fmt.Println("Goroutine: Mensagem enviada!")
	}()

	fmt.Println("Main: Esperando por mensagem...")
	mensagem := <-ch // Bloqueia aqui até a goroutine enviar
	fmt.Printf("Main: Mensagem recebida: '%s'\n", mensagem)
}
