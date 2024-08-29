package main

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"booksAdmin/internal/cli"
	"booksAdmin/internal/service"
	"booksAdmin/internal/web"
)

func main() {
	db, err := sql.Open("sqlite3", "../../books.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bookService := service.NewBookService(db)

	err = service.CreateTable(db)
	if err != nil {
		panic(err)
	}

	bookHandlers := web.NewBookHandlers(bookService)

	if len(os.Args) > 1 && (os.Args[1] == "simulate" || os.Args[1] == "search") {
		BookCLI := cli.NewBookCLI(bookService)
		BookCLI.Run()
		return
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /books", bookHandlers.GetBooks)
	router.HandleFunc("GET /books/{id}", bookHandlers.GetBookByID)
	router.HandleFunc("GET /books/search", bookHandlers.GetBookByTitle)
	router.HandleFunc("POST /books", bookHandlers.CreateBook)
	router.HandleFunc("PUT /books/{id}", bookHandlers.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandlers.DeleteBook)
	router.HandleFunc("POST /books/simulate-reading", bookHandlers.SimulateReading)

	http.ListenAndServe(":8080", router)
}
