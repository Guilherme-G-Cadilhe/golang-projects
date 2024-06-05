package main

import (
	"fmt"
	"html/template" // ou text/template
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Cursos{
			{Nome: "Golang", CargaHoraria: 40},
			{Nome: "Java", CargaHoraria: 20},
			{Nome: "Python", CargaHoraria: 5},
			{Nome: "C#", CargaHoraria: 10},
		})
		if err != nil {
			fmt.Println("error executing template:", err)
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)

}
