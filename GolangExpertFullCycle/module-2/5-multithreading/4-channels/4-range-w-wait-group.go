package main

import "sync"

func main4() {

	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	// Duas threads diferentes
	go publish2(ch)
	go reader2(ch, &wg)

	// Segura a thread principal
	wg.Wait()

}

func reader2(ch <-chan int, wg *sync.WaitGroup) {
	for n := range ch {
		println(n)
		wg.Done()
	}
}

func publish2(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // Se não fechar o canal, ele fica aberto infinitamente esperando um proximo laço do loop, que já finalizou
}
