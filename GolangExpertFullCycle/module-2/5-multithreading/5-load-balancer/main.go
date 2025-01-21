package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(workerId int, data <-chan int, wg *sync.WaitGroup) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}
func main() {
	data := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(100)
	QtdWorkers := 10
	for i := 0; i < QtdWorkers; i++ {
		go worker(i, data, &wg)
	}

	for i := 0; i < 100; i++ {
		data <- i

	}

	wg.Wait()
}
