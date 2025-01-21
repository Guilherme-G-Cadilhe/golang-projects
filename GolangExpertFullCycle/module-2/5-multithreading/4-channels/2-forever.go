package main

// Thread 1
func main2() {
	// foreverCh := make(chan bool) // Canal vazio

	// // Thread 2 funciona normal
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		println(i)
	// 	}
	// }()

	// // Thread 1 travada
	// <-foreverCh // Canal fica vazio e fica em loop infinito (deadlock)

	// =========================================================================

	foreverCh := make(chan bool) // Canal vazio

	// Canais foram feitos para se comunicar entre goroutines
	// foreverCh <- true // Não funciona, pois precisa de uma goroutine para encher o canal

	// Thread 2
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		foreverCh <- true // Enche o canal
	}()

	<-foreverCh // Canal Esvazia normalmente

}
