package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// Usamos um slice, que é mais flexível que um array.
	output [30]string
	oi     = 0
	mu     sync.Mutex // Nosso Mutex para proteger o acesso
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2) // Apenas para as duas goroutines que lançamos

	go loop("Goroutine 1", &wg)
	go loop("Goroutine 2", &wg)

	// A main também participa, mas não precisa de WaitGroup.
	loop("Main", nil)

	wg.Wait()

	for _, v := range output {
		if v != "" { // Imprime apenas os valores preenchidos
			fmt.Println(v)
		}
	}
}

func loop(name string, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	for i := 0; i < 10; i++ {
		// Antes de tocar nas variáveis globais, travamos o mutex.
		mu.Lock()
		if oi < len(output) {
			output[oi] = fmt.Sprintf("%s: %d", name, i)
			oi++
		}
		// Depois de terminar, liberamos a trava para outra goroutine usar.
		mu.Unlock()

		// Dizemos ao scheduler para dar a vez a outra goroutine forçadamente. por que números pequenos no loop fazem a tarefa ser tão rapida, que o scheduler não precisa aguardar.
		runtime.Gosched()
	}
}
