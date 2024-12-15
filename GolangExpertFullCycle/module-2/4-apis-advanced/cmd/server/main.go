package main

import "github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-2/4-apis-advanced/configs"

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(config.DBDriver)
}
