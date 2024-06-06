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

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	// Template new => Escolhe o arquivo a ser renderizado baseado no nome passado em relação ao Slice de templates
	// Dentro do Content.html => {{template "header.html"}} insere o arquivo header.html na renderização
	t := template.Must(template.New("content.html").ParseFiles(templates...))
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
