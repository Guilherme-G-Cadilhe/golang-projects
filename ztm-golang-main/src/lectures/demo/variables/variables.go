package main

// Importa o package Main e especifica que o arquivo atual é o principal, para isso basta colocar o package main.

import "fmt" // Usado para imprimir no terminal

func main() { // Função principal do programa
	var myName = "Jayson"
	fmt.Println("My Firstname is", myName)

	var name string = "Kathy"
	fmt.Println("Your Lastname is", name)

	userName := "admin"
	fmt.Println("Your username is", userName)

	var (
		firstName = "Jayson"
		lastName  = "Kathy"
		fullName = fmt.Sprintf("%s %s", firstName, lastName)
	)
	fmt.Println("Your Fullname is", fullName)

	a, b, c := 1, 2, 3
	fmt.Println("a",a, "b", b, "c", c)

	var sum int
	fmt.Println(sum)
	sum = sum + a + b + c
	fmt.Println(sum)

	part1, other := 1,5
	fmt.Println("part1 is", part1, " and other is",  other)

	part2, other := 2,0
	fmt.Println("part1 is", part2, " and other is",  other)

	word1, word2, _ := "hello", "world", "!" // _ é ignorado
	fmt.Println(word1, word2)

}
