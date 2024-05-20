package main

import "fmt"

func main() {
	// slice não declara valor fixo, diferente de array
	slice := []int{10, 20, 30, 50, 70, 100}

	fmt.Printf("O tipo do Slice é %T\n", slice)
	fmt.Printf("length=%d capacity=%d %v\n", len(slice), cap(slice), slice)

	// Slice é uma forma de definir quantos items do elemento será retornado apartir do index informado
	// slice[start:end]
	// creates an empty sub-slice of length 0 and capacity 5
	fmt.Printf("length=%d capacity=%d %v\n", len(slice[:0]), cap(slice[:0]), slice[:0])

	// creates a sub-slice with the first 4 elements
	fmt.Printf("length=%d capacity=%d %v\n", len(slice[:4]), cap(slice[:4]), slice[:4])

	// creates a sub-slice starting from the 3rd element to the end of the slice
	fmt.Printf("length=%d capacity=%d %v\n", len(slice[2:]), cap(slice[2:]), slice[2:])

	// Adiciona um elemento acima da capacidade atual, o Go copia o array e cria um novo com o dobro de capacidade
	slice = append(slice, 200)
	fmt.Printf("length=%d capacity=%d %v\n", len(slice[:2]), cap(slice[:2]), slice[:2])

}
