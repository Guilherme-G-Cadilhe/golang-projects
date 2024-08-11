package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func cancelContext() {
	// Contexto pai com cancelamento
	parentCtx, parentCancel := context.WithCancel(context.Background())

	// Contexto filho com cancelamento
	childCtx, childCancel := context.WithCancel(parentCtx)

	go func() {
		time.Sleep(2 * time.Second)
		parentCancel()
	}()

	select {
	case <-childCtx.Done():
		fmt.Println("Contexto filho cancelado por que contexto pai foi cancelado")
	case <-time.After(3 * time.Second):
		fmt.Println("Contexto filho não foi cancelado dentro do tempo limite")
	}

	childCancel()
}

func cancelWhenOneReturnTrue() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ch := make(chan bool, 10)

	get := func(ctx context.Context, ch chan bool, number int) {
		time.Sleep(4 * time.Second)
		// fazer os calculos
		if number == 1 {
			ch <- true
		}
		ch <- true
	}

	for i := 0; 1 < 10; i++ {

		go get(ctx, ch, rand.Intn(10))

	}

	select {
	case <-ctx.Done():
		fmt.Println("Contexto cancelado")
	default:
		value := <-ch
		if value {
			cancel()
		}

	}

}

type key int

const requestIDKey key = 0

func contextTestWTimeout() {
	// Configura o logger
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// Cria um contexto com timeout de 10 segundo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Assegura que o contexto será cancelado ao final

	// Associa um ID de requisição ao contexto
	ctx = context.WithValue(ctx, requestIDKey, "req-12345")

	// Grupo de espera para as goroutines
	var waitGroup sync.WaitGroup

	// Função que realiza uma requisição HTTP simulada
	requestFunc := func(ctx context.Context, url string) {
		defer waitGroup.Done()
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			log.Printf("[%s]Erro ao criar requisição: %v", getRequestID(ctx), err)
			return
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("[%s]Erro ao realizar requisição: %v", getRequestID(ctx), err)
			return
		}
		defer res.Body.Close()
		log.Printf("[%s]Erro ao realizar requisição: %v", getRequestID(ctx), err)

		log.Printf("[%s]Requisição para %s completada com status: %s", getRequestID(ctx), url, res.Status)
	}

	// Inicia várias goroutines que fazem requisições HTTP
	urls := []string{
		"https://httpbin.rog/delay/2",
		"https://httpbin.rog/delay/3",
		"https://httpbin.rog/delay/4",
	}
	for _, url := range urls {
		waitGroup.Add(1)
		go requestFunc(ctx, url)
	}

	// Aguarda que todas as requisições sejam concluídas
	waitGroup.Wait()
	log.Printf("[%s]Todas as requisições concluídas", getRequestID(ctx))
}

func getRequestID(ctx context.Context) string {
	// Na hora de pegar o valor como (any), lembrar de converter para o tipo correto (string)
	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		return requestID
	}
	return "unknown"
}

// Função que simula uma tarefa complexa que pode ser cancelada
func complexTask(ctx context.Context, result chan<- string) {
	// Usando uma goroutine para executar a tarefa
	go func() {
		select {
		case <-time.After(3 * time.Second):
			// Tarefa completa após 3 segundos
			result <- "Tarefa completa com sucesso"
		case <-ctx.Done():
			// Se o contexto for cancelado antes de completar, envia o erro pelo channel
			result <- fmt.Sprintf("Tarefa cancelada: %v", ctx.Err())
		}
	}()
}

// Função que cria um pipeline de contexto com cancelamento e timeout
func contextPipeline() {
	// Contexto de fundo, que serve como raiz do pipeline
	ctx := context.Background()

	// Primeiro nível do pipeline: contexto com valor
	ctxValue := context.WithValue(ctx, "requestID", "12345")

	// Segundo nível: contexto com timeout de 2 segundos
	ctxTimeout, cancel := context.WithTimeout(ctxValue, 2*time.Second)
	defer cancel() // Libera recursos após o uso

	// Canal para receber o resultado da tarefa
	result := make(chan string)

	// Executa a tarefa complexa com o contexto encadeado
	complexTask(ctxTimeout, result)

	// Aguarda o resultado da tarefa ou o término do contexto
	select {
	case res := <-result:
		fmt.Println(res)
	case <-ctxTimeout.Done():
		fmt.Println("Pipeline cancelado ou expirado:", ctxTimeout.Err())
	}
}

// Função que cancela um grupo de tarefas simultaneamente
func cancelMultipleTasks(ctx context.Context) {
	// Channel para sinalizar o cancelamento de todas as tarefas
	cancelChan := make(chan struct{})

	// Função auxiliar para simular uma tarefa
	runTask := func(taskID int, ctx context.Context, cancelChan <-chan struct{}) {
		select {
		case <-time.After(time.Duration(taskID) * time.Second):
			fmt.Printf("Tarefa %d completa\n", taskID)
		case <-cancelChan:
			fmt.Printf("Tarefa %d cancelada\n", taskID)
		case <-ctx.Done():
			fmt.Printf("Tarefa %d interrompida: %v\n", taskID, ctx.Err())
		}
	}

	// Executa 3 tarefas em paralelo
	for i := 1; i <= 3; i++ {
		go runTask(i, ctx, cancelChan)
	}

	// Cancela todas as tarefas após 1.5 segundos
	time.Sleep(1 * time.Second)
	close(cancelChan)

	// Aguarda a finalização das tarefas
	time.Sleep(3 * time.Second)
}

// Função que demonstra o uso de Contexto com Deadline
func contextWithDeadline() {
	// Define o deadline para 10:00 da manhã
	deadline := time.Now().Add(time.Until(time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		10, 0, 0, 0, time.Local,
	)))
	// now := time.Now()
	// deadline2 := time.Date(now.Year(), now.Month(), now.Day(), 10, 0, 0, 0, now.Location())

	// Cria um contexto com o deadline definido
	ctxDeadline, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// Canal para receber o resultado da tarefa
	result := make(chan string)

	// Executa a tarefa complexa com o contexto de deadline
	complexTask(ctxDeadline, result)

	// Aguarda o resultado da tarefa ou o término do contexto
	select {
	case res := <-result:
		fmt.Println(res)
	case <-ctxDeadline.Done():
		fmt.Println("Contexto expirado devido ao deadline:", ctxDeadline.Err())
	}
}

func main() {
	fmt.Println("Exemplo 1: Pipeline de Contexto com Encadeamento e Timeout")
	contextPipeline()

	fmt.Println("\nExemplo 2: Cancelamento de Múltiplas Tarefas Simultâneas")
	// Contexto com cancelamento manual para as tarefas
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cancelMultipleTasks(ctx)

	fmt.Println("\nExemplo 3: Contexto com Deadline")
	contextWithDeadline()
}
