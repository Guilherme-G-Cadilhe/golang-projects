package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
)

var number uint64 = 0

// PROBLEMA DE CONCORRENCIA POR QUE CADA REQUISIÇÃO HTTP CRIA NOVA THREAD QUE TENTA MEXER NO NUMBER
func main() {
	// Método  com Mutex
	mu := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()   // Bloqueia o acesso a `number`
		number++    // Incrementa o contador
		mu.Unlock() // Desbloqueia após a modificação
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essa página: %d vezes", number)))
	})

	// Método Atomic
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&number, 1) // Incrementa o contador de forma atômica
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essa página: %d vezes", number)))
	})

	http.ListenAndServe(":3000", nil)
}
