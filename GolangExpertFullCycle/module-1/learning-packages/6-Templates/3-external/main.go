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

type Cursos []Curso

func main() {

	/*
			Dentro do HTML
			- {{.}}(ponto) representa o objeto na raiz, se tiver propriedades aninhadas:
		    { teste: { propriedade1: 30, propriedade2: 40 } }
		    {{ .teste.propriedade1 }}

	*/
	t := template.Must(template.New("template.html").ParseFiles("./template.html"))

	err := t.Execute(os.Stdout, Cursos{
		{Nome: "Golang", CargaHoraria: 40},
		{Nome: "Java", CargaHoraria: 20},
		{Nome: "Python", CargaHoraria: 5},
		{Nome: "C#", CargaHoraria: 10},
	})
	if err != nil {
		fmt.Println("error executing template:", err)
		panic(err)
	}

}
