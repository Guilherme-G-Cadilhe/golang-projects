package main

import (
	"fmt"

	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-2/6-manipulando-eventos/2-integrando-rabbitmq/rabbitmq"
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
