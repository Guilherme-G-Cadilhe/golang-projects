package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Configura um servidor HTTP na porta 8080
	server := &http.Server{Addr: ":8080"}

	// Define um handler que simula uma operação demorada de 8 segundos
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(8 * time.Second)
		w.Write([]byte("Hello, World!"))
	})

	// Inicia o servidor em uma goroutine para que ele rode em segundo plano
	go func() {
		// ListenAndServe inicia o servidor HTTP e bloqueia até que ocorra um erro
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen and serve: %v\n", err)
		}
	}()

	// Cria um canal para receber sinais do sistema operacional
	stop := make(chan os.Signal, 1)
	// Notifica o canal quando receber sinais de interrupção ou terminação
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// Bloqueia a goroutine principal até que um sinal seja recebido
	<-stop

	// Cria um contexto com timeout de 5 segundos para o processo de desligamento gracioso
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Shutting down server...")

	// Tenta desligar o servidor graciosamente, esperando as requisições em andamento terminarem
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	fmt.Println("Server exiting")

	/*
		  Funcionamento do <-stop
		O que é <-stop?: Este é um operador que bloqueia a execução da goroutine atual (neste caso, a goroutine principal) até que o canal stop receba um valor.
		Por que <-stop impede a execução do restante do código?: Porque o operador <- está esperando por um valor do canal stop. Isso faz com que a execução pare e espere até que um sinal seja enviado para o canal. Somente depois que o sinal é recebido, a execução do código continua.
		Threads e Goroutines
		O servidor está em outra thread?: No Go, usamos o termo "goroutine" em vez de "thread". Goroutines são mais leves e são gerenciadas pelo runtime do Go. O servidor HTTP está sendo executado em uma goroutine separada.
		O que são goroutines?: Goroutines são funções ou métodos que são executados de forma concorrente com outras goroutines na mesma aplicação. Elas são iniciadas com a palavra-chave go.
		Por que o servidor continua funcionando?: Porque ele está em uma goroutine separada. A goroutine principal está bloqueada esperando o sinal (<-stop), mas a goroutine do servidor (go server.ListenAndServe()) continua executando em paralelo.
	*/

}
