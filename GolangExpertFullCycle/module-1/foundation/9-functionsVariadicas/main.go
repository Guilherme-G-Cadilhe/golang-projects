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

	// Chamando a função variádica com um slice desembrulhado
	numbers := []int{4, 5, 6}
	result2 := sum(numbers...)
	fmt.Println("Sum of 4, 5, 6:", result2) // Output: Sum of 4, 5, 6: 15

	// Chamando a função variádica sem argumentos
	result3 := sum()
	fmt.Println("Sum with no arguments:", result3) // Output: Sum with no arguments: 0
}
