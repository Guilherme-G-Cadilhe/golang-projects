package main

import (
	"fmt"
	"html/template" // ou text/template
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {

	curso := Curso{Nome: "Golang", CargaHoraria: 40}

	// Must => Retorna um erro se o template estiver inválido
	templateCurso := template.Must(template.New("template").Parse("O curso de {{.Nome}} tem carga horária de {{.CargaHoraria}} horas"))

	// Executa o template e o imprime na tela
	err := templateCurso.Execute(os.Stdout, curso)
	if err != nil {
		fmt.Println("error executing template:", err)
		panic(err)
	}

}
