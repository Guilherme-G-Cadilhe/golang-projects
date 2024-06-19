package main

import (
	"fmt"
)

func main() {

	// temos apenas o FOR
	for i := 0; i < 10; i++ {
		print(i)
	}

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(numeros); i++ {
		print(numeros[i])
	}

	nomes := []string{"Guilherme", "João", "Maria"}
	// i = indice e v = value, a nomeação pode ser qualquer coisa, por que ele sempre retorna 2 valores, o primeiro é sempre o index e o segundo é sempre o valor
	// Estilo do FOR IN do JS
	for i, v := range nomes {
		println(i, v, nomes[i])
	}

	// Estilo do WHILE do JS
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// LOOP INFINITO
	// for {
	// 	println("Loop Infinito")
	// }

	leetCode := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(leetCode, target)
	fmt.Println(" Result twoSums: ", result)

}
func twoSum(nums []int, target int) []int {
	// https://leetcode.com/problems/two-sum
	// Time: O(n) ( If there is many items, will look at them just once)
	// Space: O(n) ( Will add at most n items (the current length) to the map)
	result := make(map[int]int)
	for index, value := range nums {
		println("Index: ", index, " Value: ", value)
		possibleNumber := target - value

		if _, ok := result[possibleNumber]; ok {
			return []int{result[possibleNumber], index}
		}
		result[value] = index

	}
	return nil
}
