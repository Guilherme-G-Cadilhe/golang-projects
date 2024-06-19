package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("----Request started-----")
	defer log.Println("Request ended")

	select {
	// Espera 5 segundos antes de retornar uma resposta
	case <-time.After(5 * time.Second):
		log.Println("request processada com sucesso")
		// Retorna uma resposta HTTP 200
		w.Write([]byte("Request processada com sucesso"))

		// Caso o contexto seja cancelado, retorna uma resposta HTTP 408
	case <-ctx.Done():
		log.Println("request cancelada pelo cliente")
		http.Error(w, "request cancelada pelo cliente", http.StatusRequestTimeout)
	}
}
