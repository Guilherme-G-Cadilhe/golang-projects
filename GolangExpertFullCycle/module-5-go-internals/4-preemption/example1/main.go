// Este código mostra a preempção em ação.
// Temos uma goroutine "egoísta" que tenta usar 100% da CPU em um loop.
// Mesmo assim, a goroutine "educada" (na main) consegue imprimir mensagens.
// Isso prova que o scheduler está interrompendo o loop para dar a vez a outra.
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Garante que o scheduler tenha núcleos para trabalhar
	runtime.GOMAXPROCS(2)

	// Goroutine "egoísta" que entra em um loop infinito
	go func() {
		fmt.Println("🔥 Goroutine egoísta iniciada!")
		// Desafia o scheduler com um loop sem pontos de cooperação.
		for {
		}
	}()

	// A goroutine principal continua seu trabalho sem ser bloqueada.
	for i := 0; i < 10; i++ {
		fmt.Printf("✅ Principal: Tick %d\n", i)
		// time.Sleep também informa ao scheduler para executar outra coisa.
		time.Sleep(1 * time.Second)
	}

	fmt.Println("🎉 Principal terminada. O programa permaneceu responsivo!")
}

// Saída esperada:
// 🔥 Goroutine egoísta iniciada!
// ✅ Principal: Tick 0
// ✅ Principal: Tick 1
// ... (e assim por diante)
// 🎉 Principal terminada. O programa permaneceu responsivo!
