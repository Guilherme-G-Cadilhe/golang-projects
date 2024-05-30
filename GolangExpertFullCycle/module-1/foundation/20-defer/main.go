package main

import "fmt"

func main() {

	fmt.Println("Primeira Linha")

	// O defer é executado após o final do bloco de código
	defer fmt.Println("Segunda Linha")

	fmt.Println("Terceira linha")

	/*
	   > Primeira linha
	   > Terceira linha
	   > Segunda linha
	*/

}
