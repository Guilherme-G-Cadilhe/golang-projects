package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0 // Usando atomic para não ter lock

	// RabbitMQ
	go func() {
		// time.Sleep(1 * time.Second)
		// msg := Message{1, "Hello from RabbitMQ"}
		// c1 <- msg
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(1 * time.Second)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	// Kafka
	go func() {
		// time.Sleep(2 * time.Second)
		// msg := Message{1, "Hello from Kafka"}
		// c2 <- msg
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(2 * time.Second)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()

	// // seleciona a primeira mensagem recebida ou um timeout se passar 3 segundos
	// select {
	// case msg1 := <-c1:
	// 	println("received", msg1)

	// case msg2 := <-c2:
	// 	println("received", msg2)

	// case <-time.After(3 * time.Second):
	// 	println("timeout")

	// default: // primeira coisa a acontecer
	// 	println("no message received")
	// }

	// Cria um loop para executar o select 3 vezes
	// for i := 0; i < 3; i++ {
	// Ou Loop infinito para estar sempre escutando
	for {
		select {
		case msg := <-c1: // ex: rabbitMq
			fmt.Printf("received from rabbit %d: %s\n", msg.id, msg.Msg)

		case msg := <-c2: // ex: kafka
			fmt.Printf("received from kafka %d: %s\n", msg.id, msg.Msg)

		case <-time.After(3 * time.Second):
			println("timeout")

			// default: // primeira coisa a acontecer se nenhum canal estiver pronto ainda
			// 	println("no message received")
		}
	}

}
