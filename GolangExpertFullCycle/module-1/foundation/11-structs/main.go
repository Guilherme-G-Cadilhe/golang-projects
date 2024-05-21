package main

import "fmt"

// Structs é como se fosse um tipo de uma Classe
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	guilherme := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	fmt.Println(guilherme)
	guilherme.Ativo = false
	fmt.Printf("Nome: %s Idade: %d Ativo: %t\n", guilherme.Nome, guilherme.Idade, guilherme.Ativo)

}
