package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {
	// Definindo um limite de memoria de 100MB
	// debug.SetMemoryLimit(100 * 1024 * 1024) // 100 MB
	debug.SetMemoryLimit(10 * 1024 * 1024) // 10 MB

	// Função para alocar memoria
	allocateMemory := func(size int) []byte {
		data := make([]byte, size)
		return data
	}

	// Alocando memoria para observar o comportamento do GC
	for i := 0; i < 10; i++ {
		_ = allocateMemory(20 * 1024 * 1024) //Aloca 20 MB de memoria
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc = %v MB\n", m.Alloc/1024/1024)
		fmt.Printf("TotalAlloc = %v MB\n", m.TotalAlloc/1024/1024)
		fmt.Printf("Sys = %v MB\n", m.Sys/1024/1024)
		fmt.Printf("Lookups = %v\n", m.Lookups)
		fmt.Printf("Mallocs = %v\n", m.Mallocs)
		fmt.Printf("Frees = %v\n", m.Frees)
		fmt.Printf("HeapAlloc = %v MB\n", m.HeapAlloc/1024/1024)
		fmt.Printf("HeapSys = %v MB\n", m.HeapSys/1024/1024)
		fmt.Printf("HeapIdle = %v MB\n", m.HeapIdle/1024/1024)
		fmt.Printf("HeapInuse = %v MB\n", m.HeapInuse/1024/1024)
		fmt.Printf("HeapReleased = %v MB\n", m.HeapReleased/1024/1024)
		fmt.Printf("HeapObjects = %v\n", m.HeapObjects)
		fmt.Printf("StackInuse = %v MB\n", m.StackInuse/1024/1024)
		fmt.Printf("StackSys = %v MB\n", m.StackSys/1024/1024)
		fmt.Printf("MSpanInuse = %v MB\n", m.MSpanInuse/1024/1024)
		fmt.Printf("MSpanSys = %v MB\n", m.MSpanSys/1024/1024)
		fmt.Printf("MCacheInuse = %v MB\n", m.MCacheInuse/1024/1024)
		fmt.Printf("MCacheSys = %v MB\n", m.MCacheSys/1024/1024)
		fmt.Printf("BuckHashSys = %v MB\n", m.BuckHashSys/1024/1024)
		fmt.Printf("GCSys = %v MB\n", m.GCSys/1024/1024)
		fmt.Printf("OtherSys = %v MB\n", m.OtherSys/1024/1024)
		fmt.Printf("NumGC = %v\n", m.NumGC)
	}

	// Exibindo o uso de memoria

	// Rodar
	// GODEBUG=gctrace=1 go run main.go

}
