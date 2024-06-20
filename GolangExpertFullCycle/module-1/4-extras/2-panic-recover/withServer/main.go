package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

// Middleware que utiliza recover para capturar panics
func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// Captura qualquer panic que ocorrer
			if err := recover(); err != nil {
				// Loga o panic e a stack trace
				fmt.Println("Recovered panic:", err)
				debug.PrintStack()
				// Retorna um erro interno do servidor ao cliente
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		// Continua com o próximo handler na cadeia
		fmt.Println("Recover middleware")
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Multiplexador de requests
	mux := http.NewServeMux()
	// Handler para a rota raiz
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	// Handler que provoca um panic
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Panic!")
	})

	log.Println("Listening on port 8080")
	// Inicia o servidor HTTP com o middleware de recover
	if err := http.ListenAndServe(":8080", recoverMiddleware(mux)); err != nil {
		log.Fatalf("Could not listen and serve: %v\n", err)
	}
}
