package main

import (
	"fmt"
	"teste/rabbitmq"
)

func main() {

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	for i := 0; i < 10; i++ {
		err = rabbitmq.Publish(ch, fmt.Sprintf("Mensagem %d", i), "amq.direct")
		if err != nil {
			panic(err)
		}
	}
}
