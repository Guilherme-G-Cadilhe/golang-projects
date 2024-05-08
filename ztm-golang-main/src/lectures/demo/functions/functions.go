package main

import "fmt"

func sum(leftHand, rightHand int) int {
	return leftHand + rightHand
}

func multiReturn() (int, int, int) {
	return 1, 2, 3
}

func double(x int) int {
	return x * 2
}
func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello, World!")
}

func main() {

	result := sum(1, 2)
	fmt.Println(result)

	a, b , _ := multiReturn() // _ é ignorado
	fmt.Println(a, b)

	result = double(6)
	fmt.Println(result)

	result = add(result, 2)
	fmt.Println(result)

	greet()



}
