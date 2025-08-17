package main

import (
	"fmt"
	"runtime"
)

var (
	output [30]string // 3 times, 10 iterations each
	oi     = 0
)

// RODANDO GOROUTINES PARALELAS UMA POR VEZ
func main() {

	runtime.GOMAXPROCS(1)
	chanFinished1 := make(chan bool)
	chanFinished2 := make(chan bool)

	go loop("Goroutine 1", chanFinished1)
	go loop("Goroutine 2", chanFinished2)
	loop("Main", nil)

	<-chanFinished1
	<-chanFinished2

	for _, v := range output {
		fmt.Println(v)
	}

}

func loop(name string, finished chan bool) {
	for i := 0; i < 10; i++ {
		output[oi] = fmt.Sprintf("%s: %d", name, i)
		oi++
	}

	if finished != nil {
		finished <- true
	}
}
