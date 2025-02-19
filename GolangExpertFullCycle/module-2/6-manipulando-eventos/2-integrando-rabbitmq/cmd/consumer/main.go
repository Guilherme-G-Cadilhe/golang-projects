package main

import (
	"fmt"
	"teste/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)        // canal de mensagens
	go rabbitmq.Consume(ch, msgs, "orders") // consumindo as mensagens

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false) // ACK e o false é para não dar ack em todas as outras msgs da fila junto
	}
}
