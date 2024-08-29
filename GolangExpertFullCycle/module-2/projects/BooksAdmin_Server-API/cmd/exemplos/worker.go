package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

// ---- Exemplo 2 -----
func main() { // thread goroutine 1
	ch := make(chan int) // empty

	qtdWorkers := 10

	// Call the workers
	for i := range qtdWorkers {
		go worker(i, ch)
	}

	for i := range 10 {
		ch <- i
	}
}

// ---- Exemplo 1 -----
// func main() { // thread goroutine 1
// 	ch := make(chan string) // empty

// 	// goroutine 2
// 	go func() {
// 		ch <- "Full Cycle"
// 	}()

// 	msg := <-ch

// 	fmt.Println(msg)
// }
