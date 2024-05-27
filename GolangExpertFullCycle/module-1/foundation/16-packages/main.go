package main

// import matematica "packageExample/matematica"

import (
	"packageExample/matematica"

	"github.com/google/uuid"
)

func main() {

	soma := matematica.Soma(10, 20)
	println(soma)
	carro := matematica.Carro{
		Marca: "Fiat",
	}
	println(carro.Marca)

	println(matematica.A)

	// Função 'ligar' é undefined por que não está sendo exportada com inicial Maiuscula
	// println(carro.ligar())

	println(uuid.New().String())

}
