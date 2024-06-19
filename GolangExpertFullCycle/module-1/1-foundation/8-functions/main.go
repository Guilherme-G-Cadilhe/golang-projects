package main

import (
	"errors"
	"fmt"
)

func main() {

	result := sum(10, 20)
	fmt.Println(result)

	result, even := doubleAndEven(10)
	fmt.Println(result, even)

	a, err := funcWithError("hello world")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}

}

func sum(a, b int) int {
	return a + b
}
func doubleAndEven(x int) (int, bool) {
	result := x * 2
	return result, result%2 == 0
}

func funcWithError(a string) (string, error) {

	if len(a) > 5 {
		return "", errors.New("string too long")
	}
	return a, nil
}
