package main

import "net/http"

func main() {
	// request multiplexer for HTTP servers in Go that can be used to register handlers for different HTTP methods
	mux := http.NewServeMux()

	// mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	// 	res.Write([]byte("Hello, World!"))
	// })
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "Blog de Legumes"})

	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, Guilherme!"))
	})
	http.ListenAndServe(":8081", mux2)
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("CHEGOU NA HOME!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(b.title))
}
