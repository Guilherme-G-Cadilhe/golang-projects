package main

import (
	"log"
	"net/http"
)

func main() {
	// Handler que disponibiliza arquivos estáticos no diretório 'public'
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()

	// Disponibiliza index.html para a rota '/'
	mux.Handle("/", fileServer)

	// Disponibiliza pra Web uma API com o caminho /api/blog
	mux.HandleFunc("/api/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Log fatal para encerrar o servidor se houver um erro e imprimir no terminal
	log.Fatal(http.ListenAndServe(":8080", mux))

}
