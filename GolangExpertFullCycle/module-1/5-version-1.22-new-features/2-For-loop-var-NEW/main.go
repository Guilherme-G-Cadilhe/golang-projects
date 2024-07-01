package main

import "fmt"

func main() {
	// Old extensive form
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// New form in Go 1.22
	for i := range 5 {
		fmt.Println(i)
	}

	done := make(chan bool)
	values := []string{"Guilherme", "João", "Maria"}
	for _, v := range values {
		// No need for v := v workaround in Go 1.22
		go func(val string) {
			fmt.Println(val)
			done <- true
		}(v)
	}
	for range values {
		<-done
	}
}
