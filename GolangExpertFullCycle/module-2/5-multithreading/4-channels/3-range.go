package main

func main3() {

	ch := make(chan int)
	go publish(ch)
	reader(ch) // sem ser goroutine para que a thread 1 não finalize precocemente

	// em outra func ( reader) ou direto
	// for n := range ch {
	// 	println(n)
	// }
}

func reader(ch <-chan int) {
	for n := range ch {
		println(n)
	}
}

func publish(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // Se não fechar o canal, ele fica aberto infinitamente esperando um proximo laço do loop, que já finalizou
}
