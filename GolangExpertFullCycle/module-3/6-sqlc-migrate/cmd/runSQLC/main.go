package main

import (
	"context"
	"database/sql"

	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-3/6-sqlc-migrate/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)
	err = dbConn.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	// ==== FLUXO 1: CRIA CATEGORIA ====

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:   uuid.New().String(),
	// 	Name: "Golang",
	// 	Description: sql.NullString{
	// 		String: "Curso de Golang",
	// 		Valid:  true,
	// 	},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description.String)
	// }

	// ==== FLUXO 2: ATUALIZA CATEGORIA ====

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:   "ce2cf10f-c16b-4772-a9ac-a9ff3bdd716d",
	// 	Name: "Go Updated",
	// 	Description: sql.NullString{
	// 		String: "Curso de Golang Updated",
	// 		Valid:  true,
	// 	},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description.String)
	// }

	// ==== FLUXO 3: DELETA CATEGORIA ====

	err = queries.DeleteCategory(ctx, "ce2cf10f-c16b-4772-a9ac-a9ff3bdd716d")
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}
