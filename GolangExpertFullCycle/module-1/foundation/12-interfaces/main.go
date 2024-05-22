package main

import "fmt"

// Definindo a interface Speaker
type Speaker interface {
	Speak() string
}

// Tipo Person implementando a interface Speaker
type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

// Tipo Dog implementando a interface Speaker
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof! My name is " + d.Name
}

// Função que aceita qualquer tipo que implemente Speaker
func Greet(s Speaker) {
	fmt.Println(s.Speak())
}

// Função que aceita qualquer tipo usando a interface vazia
func PrintAnything(a interface{}) {
	fmt.Println(a)
}

func showType(a interface{}) {
	fmt.Printf("Type: %T, e o valor: %v\n", a, a)
}

func main() {
	p := Person{Name: "Alice"}
	d := Dog{Name: "Buddy"}

	Greet(p) // Output: Hello, my name is Alice
	Greet(d) // Output: Woof! My name is Buddy

	// Usando a interface vazia
	PrintAnything(42)
	PrintAnything("Hello")
	PrintAnything(p)
	var x interface{} = 42
	var y interface{} = "Hello"
	showType(x)
	showType(y)
}
