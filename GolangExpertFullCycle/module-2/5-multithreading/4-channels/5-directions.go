package main

import "fmt"

// Tipagem de canais
// Seta do lado direito = Canal só vai receber
// Seta do lado esquerdo = Canal só vai esvaziar

// Seta = chan<-
func recebe(nome string, ch chan<- string) {

	ch <- nome
}

// Seta = <-chan
func ler(data <-chan string) {
	fmt.Println(<-data) // <-data
}

func main() {
	channel := make(chan string)

	go recebe("Hello World", channel)

	ler(channel)
}
