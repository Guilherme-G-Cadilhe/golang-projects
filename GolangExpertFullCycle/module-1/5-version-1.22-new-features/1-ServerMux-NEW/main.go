package main

import (
	"fmt"
	"net/http"
)

// ATUALIZAÇÕES FEITAS NA VERSÃO 1.2.2
func main() {
	// Criação de um novo ServeMux para gerenciar rotas
	mux := http.NewServeMux()

	// Define rotas e handlers para diferentes padrões de URL e métodos HTTP

	// Captura um valor dinâmico {id} na rota e chama GetBookHandler
	// mux.HandleFunc("GET /books/{id}", GetBookHandler)

	// Captura múltiplos segmentos do path após /dir/ e chama BooksPathHandler
	// mux.HandleFunc("GET /books/dir/{d...}", BooksPathHandler)

	// Define uma rota exata que bloqueia sub-rotas e chama BooksHandler
	// mux.HandleFunc("GET /books/{$}", BooksHandler)

	// Define rotas específicas com precedência
	// mux.HandleFunc("GET /books/precedence/latest", BooksPrecedenceHandler)
	// mux.HandleFunc("GET /books/precedence/{x}", BooksPrecedence2Handler)

	// Define duas rotas com possíveis conflitos (isso pode causar erro)
	mux.HandleFunc("GET /books/{s}", BooksPrecedenceHandler)
	mux.HandleFunc("GET /{s}/latest", BooksPrecedence2Handler)

	// Inicia o servidor HTTP na porta 8080
	http.ListenAndServe(":8080", mux)
}

// Handler para a rota /books/{id}
func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	// Obtém o valor dinâmico {id} da rota
	id := r.PathValue("id")
	w.Write([]byte("Book " + id))
}

// Handler para a rota /books/dir/{d...}
func BooksPathHandler(w http.ResponseWriter, r *http.Request) {
	// Obtém o valor capturado do path após /dir/
	dirpath := r.PathValue("d")
	fmt.Fprintf(w, "Directory: %s", dirpath)
}

// Handler para a rota exata /books/
func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books"))
}

// Handler para a rota /books/precedence/latest
func BooksPrecedenceHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books Precedence"))
}

// Handler para a rota /{s}/latest
func BooksPrecedence2Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books Precedence 2"))
}
