package main

import (
	"fmt"
	"sync"
	"time"
)

func main2() {
	// WaitGroup é usado para esperar que todas as goroutines terminem
	var wg sync.WaitGroup

	// Lança 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Adiciona 1 ao contador do WaitGroup

		go func(id int) {
			defer wg.Done() // Decrementa o contador quando a goroutine terminar

			for j := 0; j < 5; j++ {
				fmt.Printf("Eu sou a goroutine %d, impressão %d\n", id, j)
				time.Sleep(10 * time.Millisecond) // Pequena pausa para ver a alternância
			}
		}(i)
	}

	// Espera até que o contador do WaitGroup chegue a zero
	wg.Wait()
	fmt.Println("Todas as goroutines terminaram.")
}
