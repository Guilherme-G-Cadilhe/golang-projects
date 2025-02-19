package main

import (
	"fmt"

	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-2/6-manipulando-eventos/2-integrando-rabbitmq/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
