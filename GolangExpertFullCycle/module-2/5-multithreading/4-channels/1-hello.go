package main

import "fmt"

// Thread 1
func main1() {

	channel := make(chan string) // Canal vazio

	// Thread 2
	go func() {
		channel <- "Hello World"  // Canal Cheio
		channel <- "Hello World2" // Nâo pode colocar mais sem esvaziar
	}()

	// Thread 1
	x := <-channel // Canal Esvazia
	fmt.Println(x)
}
