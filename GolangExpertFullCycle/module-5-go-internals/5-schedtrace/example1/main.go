package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

// contador para manter o numero de goroutines vazadas

var leakedGoroutines int32

func leakGoroutine() {
	atomic.AddInt32(&leakedGoroutines, 1)
	for {
		// Loop infinito representando o vazamento
	}
}

func leakHandler(w http.ResponseWriter, r *http.Request) {
	go leakGoroutine() // Lançamos uma goroutine vazada
	fmt.Fprintf(w, "Goroutines vazada!\n")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	// Imprimimos o numero de goroutines vazadas
	fmt.Fprintf(w, "Goroutines vazadas: %d\n", atomic.LoadInt32(&leakedGoroutines))
}

func main() {
	http.HandleFunc("/leak", leakHandler)     // Endpoint para causar o vazamento
	http.HandleFunc("/status", statusHandler) // Endpoint para obter o status

	fmt.Println("Servidor iniciado em http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}

// GOMAXPROCS=1 GODEBUG=schedtrace=1 go run main.go

// go-wrk -c 20 -d 10 http://localhost:8080/leak
/*
SCHED 277ms: gomaxprocs=1 idleprocs=0 threads=5 spinningthreads=0 needspinning=1 idlethreads=1 runqueue=0 [1]
SCHED 278ms: gomaxprocs=1 idleprocs=0 threads=5 spinningthreads=0 needspinning=1 idlethreads=1 runqueue=0 [1]
SCHED 280ms: gomaxprocs=1 idleprocs=0 threads=5 spinningthreads=0 needspinning=1 idlethreads=1 runqueue=0 [1]
SCHED 327ms: gomaxprocs=1 idleprocs=1 threads=5 spinningthreads=0 needspinning=0 idlethreads=1 runqueue=0 [0]
SCHED 328ms: gomaxprocs=1 idleprocs=1 threads=5 spinningthreads=0 needspinning=0 idlethreads=1 runqueue=0 [0]
SCHED 329ms: gomaxprocs=1 idleprocs=1 threads=5 spinningthreads=0 needspinning=0 idlethreads=1 runqueue=0 [0]
SCHED 330ms: gomaxprocs=1 idleprocs=0 threads=5 spinningthreads=0 needspinning=0 idlethreads=1 runqueue=0 [0]
SCHED 332ms: gomaxprocs=1 idleprocs=0 threads=5 spinningthreads=0 needspinning=0 idlethreads=1 runqueue=0 [0]

.....

SCHED 418ms: gomaxprocs=1 idleprocs=0 threads=6 spinningthreads=0 needspinning=1 idlethreads=2 runqueue=50 [8]
SCHED 460ms: gomaxprocs=1 idleprocs=0 threads=6 spinningthreads=0 needspinning=1 idlethreads=2 runqueue=55 [9]
*/
