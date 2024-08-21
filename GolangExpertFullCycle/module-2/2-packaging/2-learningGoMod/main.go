package main

// Rodar 'go mod tidy' para atualizar o 'go.mod' e 'go.sum' adicionando biblioteca uuid ou removendo se não estiver mais sendo usada.
// go.mod e go.sum funcionam como se fosse o package.json e o package.lock.
import (
	"github.com/google/uuid"
)

// comando 'got get github.com/google/uuid' para instalar a biblioteca uuid manualmente.

func main() {

	println(uuid.New().String())

}
