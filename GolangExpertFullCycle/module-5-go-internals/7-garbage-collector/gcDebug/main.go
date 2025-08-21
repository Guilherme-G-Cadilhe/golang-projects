package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Habilitar o tracing do GC
	// debug.SetGCPercent(-1) //Desativa o GC automatico
	// Ajustar o percentual do GC (Por exemplo, para 300%)
	// debug.SetGCPercent(300)

	// Função para alocar memoria
	allocateMemory := func(size int) []byte {
		data := make([]byte, size)
		return data
	}

	// Alocando memoria para observar o comportamento do GC
	for i := 0; i < 10; i++ {
		allocateMemory(10 * 1024 * 1024) //Aloca 10 MB de memoria
		time.Sleep(1 * time.Second)
	}

	// Exibindo o uso de memoria
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MB\n", m.Alloc/1024/1024)
	fmt.Printf("TotalAlloc = %v MB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("Sys = %v MB\n", m.Sys/1024/1024)
	fmt.Printf("NumGC = %v\n", m.NumGC)

	// Rodar
	// GODEBUG=gctrace=1 go run main.go
	// GODEBUG=gctrace=1 GOGC=300 go run main.go

}
