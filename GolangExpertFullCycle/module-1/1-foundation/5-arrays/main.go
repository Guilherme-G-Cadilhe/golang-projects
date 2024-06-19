package main

import "fmt"

func main() {
	var meuArray [5]int
	meuArray[0] = 2
	meuArray[1] = 5
	meuArray[2] = meuArray[0] * meuArray[1]

	fmt.Println(meuArray) // Loga o array
	// Acessa o último elemento do array
	fmt.Println(meuArray[len(meuArray)-1])

	// i = index , v = valor, range = length do array
	for i, v := range meuArray {
		fmt.Printf("Posição: %d Valor: %d\n", i, v)
	}
}
