package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./teste.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	usecase := NewUseCase(db)

	// Use the product usecase
	product, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}
	println(product.Name)
}
