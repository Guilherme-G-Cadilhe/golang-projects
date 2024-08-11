package main

import (
	"context"
	"fmt"
	"time"
)

// Função simples que simula uma operação que leva tempo.
func doSomething(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		// A operação leva 2 segundos para completar.
		fmt.Println("Operação completada com sucesso")
	case <-ctx.Done():
		// Se o contexto for cancelado antes de completar, retorna o motivo.
		fmt.Println("Operação cancelada:", ctx.Err())
	}
}

func main() {
	// Contexto básico
	ctx := context.Background()

	fmt.Println("Exemplo 1: Contexto básico sem cancelamento")
	doSomething(ctx)

	fmt.Println("Exemplo 2: Contexto com timeout")
	// Contexto com timeout de 1 segundo
	ctxTimeout, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel() // Certifique-se de liberar os recursos.

	doSomething(ctxTimeout)

	fmt.Println("Exemplo 3: Contexto com cancelamento manual")
	// Contexto que pode ser cancelado manualmente
	ctxCancel, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(500 * time.Millisecond)
		cancel() // Cancela o contexto após 500ms
	}()

	doSomething(ctxCancel)

	fmt.Println("Exemplo 4: Passando valores no contexto")
	// Contexto que carrega um valor
	ctxValue := context.WithValue(ctx, "chave", "valor importante")

	value := ctxValue.Value("chave").(string)
	fmt.Println("Valor do contexto:", value)

	fmt.Println("Exemplo 5: Contexto encadeado")
	// Encadeando contextos
	ctxParent := context.WithValue(ctx, "parent", "valor do pai")
	ctxChild := context.WithValue(ctxParent, "child", "valor do filho")

	fmt.Println("Valor do contexto pai:", ctxChild.Value("parent"))
	fmt.Println("Valor do contexto filho:", ctxChild.Value("child"))
}
