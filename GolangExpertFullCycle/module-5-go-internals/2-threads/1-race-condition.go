// Este exemplo mostra uma race condition na prática e como corrigi-la.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var contador int64
	var wg sync.WaitGroup

	// --- Versão com RACE CONDITION ---
	// Lançamos 1000 goroutines, cada uma incrementando o contador 100 vezes.
	// O resultado esperado é 100.000, mas o resultado real será menor.
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				contador++ // Perigoso! Leitura, incremento e escrita não são uma operação única.
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Resultado com Race Condition: %d (Incorreto)\n", contador)

	// --- Versão CORRIGIDA com Atomic ---
	// Resetamos o contador para fazer o teste correto.
	contador = 0
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				// atomic.AddInt64 garante que a operação de soma seja atômica,
				// ou seja, indivisível e segura contra concorrência.
				atomic.AddInt64(&contador, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Resultado corrigido com Atomic: %d (Correto)\n", contador)
}
