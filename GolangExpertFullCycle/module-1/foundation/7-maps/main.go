package main

import "fmt"

func main() {

	salarios := map[string]int{
		"Guilherme": 6000,
		"João":     5000,
		"Maria":     3000,
	}
	// Função make é usada para inicializar e alocar memória para três tipos de dados: slices, maps e channels.
	mapMake := make(map[int]string)
	mapMake[5] = "Guilherme"
	fmt.Println(mapMake)
	mapMake2 := map[string]int{}
	mapMake2["Teste"] = 5
	fmt.Println(mapMake2)

	fmt.Println(salarios)
	fmt.Println(salarios["Guilherme"])

	delete(salarios, "Maria")
	fmt.Println("Deletando Maria: ", salarios)

	// Usando make com slice
	s := make([]int, 3, 5)
	fmt.Println(s) // Output: [0 0 0]

	// Usando make com map
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m) // Output: map[a:1 b:2]

	for nome, salario := range salarios {
		fmt.Printf("Nome: %s - Salario: %d\n", nome, salario)
	}
	// _ ignora variavel
	for _, salario := range salarios {
		fmt.Printf("Salario: %d\n", salario)
	}

}
