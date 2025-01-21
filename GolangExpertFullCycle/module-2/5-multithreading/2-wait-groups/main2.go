package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Garante que Done será chamado ao final.
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main2() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)         // Aumenta o contador para cada worker.
		go worker(i, &wg) // Inicia cada worker como uma goroutine.
	}
	wg.Wait() // Aguarda até que todas as goroutines terminem.
	fmt.Println("All workers done")
}
