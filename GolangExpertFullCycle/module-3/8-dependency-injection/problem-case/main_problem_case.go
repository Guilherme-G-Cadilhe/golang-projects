package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-3/8-dependency-injection/problem-case/product"
)

func main2() {
	db, err := sql.Open("sqlite3", "./teste.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a new product repository
	productRepository := product.NewProductRepository(db)

	// Create a new product usecase
	productUsecase := product.NewProductUsecaseProblem(productRepository)

	// Use the product usecase
	product, err := productUsecase.GetProductProblem(1)
	if err != nil {
		panic(err)
	}
	println(product.Name)
}
