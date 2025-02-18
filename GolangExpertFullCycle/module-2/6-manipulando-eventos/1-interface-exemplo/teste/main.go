package main

import "fmt"

func main() {
	evento := []string{"teste", "teste2", "teste3", "teste4"}
	// evento = evento[:2] // 0, 1 = Do primeiro (:), pega até 2 depois
	// evento = evento[2:] // 2, 3 = Pula 2 e pega todo o resto
	// evento = evento[:0] // Do primeiro (:), pega 0
	evento = append(evento[:0], evento[1:]...) // Ignora o primeiro e pega todo o resto

	fmt.Println(evento)
}
