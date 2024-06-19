package main

import "fmt"

// https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures

// Lista todos os sistemas operacionais e arquiteturas
// go tool dist list

// Listar quais são as arquiteturas do sistema operacional atual
// go env GOOS GOARCH

// Básico compilação
// go build main.go

// Compilação com go.mod
// go build ( Se tiver go.mod, que fala automaticamente com o main.go )
// > nome-modulo.exe

// Compilação com nome customizado
// go build -o nome-diferente.exe
// > nome-diferente.exe

// Lista de comandos para compilar
// GOOS=linux GOARCH=amd64 go build main.go
// GOOS=linux go build main.go

// GOOS=windows GOARCH=amd64 go build main.go
// GOOS=windows go build main.go

// GOOS=darwin GOARCH=amd64 go build main.go
// GOOS=darwin go build main.go

// BASH Executando:
// ./main.exe

// CMD Executando:
// cmd /k main.exe

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Press 'Enter' to exit...")
	// Mantém o terminal aberto para visualização
	fmt.Scanln() // Espera pela entrada do usuário
}
