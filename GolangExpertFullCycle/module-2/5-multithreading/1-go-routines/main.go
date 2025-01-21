package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1 (main) e Thread 2 (Garbage Collector)
func main() {
	// Exemplo 1 de goroutine usando a funcao task para executar as tarefas
	// Termina toda a Task A e depois termina a Task B
	// task("A")
	// task("B")

	// Exemplo 2 de goroutine usando a funcao task para executar as tarefas
	// Executa todas as tarefas em paralelo
	go task("C") // Thread 3
	go task("D") // Thread 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task Anonymous is running\n", i)
			time.Sleep(1 * time.Second)
		}
	}()
	// Nada aqui.
	// Sair
	// Por ser em paralelo, o Main não espera as tarefas terminarem a execução e após temrinar de ler as linhas e não ter nada no final, ele sai da funcao main e finaliza as tarefas sem esperar o final.
	time.Sleep(15 * time.Second) // Ocupa a Thread main para esperar o fim das tarefas
}
