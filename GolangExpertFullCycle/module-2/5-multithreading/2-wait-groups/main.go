package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		waitGroup.Done() // Decrementa o WaitGroup
	}
}

func main() {
	waitGroup := sync.WaitGroup{} // Cria um WaitGroup
	waitGroup.Add(25)             // Adiciona 25 operações previstas de GoRoutine
	go task("C", &waitGroup)      // + 10 Operações
	go task("D", &waitGroup)      // + 10 Operações
	go func() {                   // + 5 Operações
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task Anonymous is running\n", i)
			time.Sleep(1 * time.Second)
			waitGroup.Done() // Decrementa o WaitGroup
		}
	}()

	waitGroup.Wait() // Aguarda todas as operações serem concluidas
}
