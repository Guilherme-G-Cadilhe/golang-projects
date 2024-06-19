package main

import fmt "fmt"

const a = "Hello world"

type ID int
type Currency float64
type Name string
type Total int

var (
	b bool     = true  // Padrão inferido de bool é sempre false
	c Total    = 10    // Padrão 0
	d Name     = "Gui" // ""
	e Currency = 1.2   // +0.0000000c+000
	f ID       = 1
)

func main() {
	fmt.Printf("O tipo de E é %T", e) // Printa o tipo da variavel
	fmt.Printf("O tipo de E é %v", e) // Printa o valor da variavel
}
