//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-3/8-dependency-injection/problem-case/product"
	"github.com/google/wire"
)

// Set de dependências
var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	// Sempre que o product.ProductRepositoryInterface for solicitado, ele irá criar um product.ProductRepository
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewUseCase(db *sql.DB) *product.ProductUsecase {
	wire.Build(
		setRepositoryDependency,
		product.NewProductUsecase,
	)
	return &product.ProductUsecase{}
}
