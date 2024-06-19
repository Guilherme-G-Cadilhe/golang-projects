package main

import "fmt"

// Função variádica para calcular a soma de múltiplos inteiros
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// Chamando a função variádica com múltiplos argumentos
	result1 := sum(1, 2, 3)
	fmt.Println("Sum of 1, 2, 3:", result1) // Output: Sum of 1, 2, 3: 6

	// Closure é uma função que é chamada dentro de outra função
	total := func() int {
		return sum(4, 5, 6) * 2
	}()
	fmt.Println(total)

}
