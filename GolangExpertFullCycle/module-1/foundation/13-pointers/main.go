package main

import "fmt"

type Pessoa struct {
	Nome string
}

// Função linkada ao type Pessoa
func (p *Pessoa) TrocaNome(newName string) {
	p.Nome = newName
}

// Função Publica
func TrocaNome(p *Pessoa, newName string) {
	p.Nome = newName
}

// Função para modificar o valor através de um ponteiro
func increment(x *int) {
	*x = *x + 1
}

func main() {
	// Variável -> Ponteiro que tem o endereço alocado na memória do pc -> Valor
	a := 42
	fmt.Println("Initial value:", a) // Output: Initial value: 42

	p := &a
	fmt.Println("Pointer value:", p)      // Output: Pointer value: <memory_address>
	fmt.Println("Value via pointer:", *p) // Output: Value via pointer: 42

	increment(p)
	fmt.Println("Modified value:", a) // Output: Modified value: 43

	var ponteiro *int = p
	fmt.Println(ponteiro)
	*ponteiro = 100
	fmt.Println(*ponteiro)
	fmt.Println("Modified after 3rd variable:", *p) // Output: Value via pointer: 100

	guilherme := Pessoa{
		Nome: "Guilherme",
	}

	fmt.Println(guilherme)
	// Passando o ponteiro como argumento da função
	TrocaNome(&guilherme, "Joaquim")
	fmt.Println(guilherme)
	guilherme.TrocaNome("Maria")
	fmt.Println(guilherme)
}
