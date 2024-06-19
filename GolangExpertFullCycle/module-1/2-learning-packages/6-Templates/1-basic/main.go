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
	fmt.Println(curso)

	// Cria um novo template
	templateCurso := template.New("CursoTemplate")
	// Adiciona a regra de variaveis dinamicas e o body do template
	templateCurso, err := templateCurso.Parse("O curso de {{.Nome}} tem carga horária de {{.CargaHoraria}} horas")
	if err != nil {
		fmt.Println("error Parsing template:", err)
		panic(err)
	}
	// Executa o template e o imprime na tela
	err = templateCurso.Execute(os.Stdout, curso)
	if err != nil {
		fmt.Println("error executing template:", err)
		panic(err)
	}

}
