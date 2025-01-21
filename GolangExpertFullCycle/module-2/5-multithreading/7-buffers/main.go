package main

func main() {
	ch := make(chan string, 2) // Buffer = Aumenta o limite para 2 elementos no canal
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
