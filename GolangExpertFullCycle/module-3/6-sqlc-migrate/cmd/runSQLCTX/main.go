package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-3/6-sqlc-migrate/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

// FUNÇÃO CUSTOMIZADA PARA TRANSAÇÕES
func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	queries := db.New(tx)

	// Executa a função anonima que faz as queries do SQLC
	err = fn(queries)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("error on rollback: %v, original error: %v", rbErr, err)
		}
		return err
	}
	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(queries *db.Queries) error {
		var err error

		err = queries.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})

		if err != nil {
			return err
		}

		err = queries.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCategory.ID,
			Price:       argsCourse.Price,
		})

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	err = dbConn.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	// ==== FLUXO 1: CRIA CATEGORIA E CURSO ====
	// courseArgs := CourseParams{
	// 	ID:   uuid.New().String(),
	// 	Name: "Golang Course Transaction",
	// 	Description: sql.NullString{
	// 		String: "MySQL Transaction",
	// 		Valid:  true,
	// 	},
	// 	Price: 10.95,
	// }

	// categoryArgs := CategoryParams{
	// 	ID:   uuid.New().String(),
	// 	Name: "Golang Category Transaction",
	// 	Description: sql.NullString{
	// 		String: "MySQL Transaction",
	// 		Valid:  true,
	// 	},
	// }

	// courseDB := NewCourseDB(dbConn)

	// err = courseDB.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)
	// if err != nil {
	// 	panic(err)
	// }

	// ==== FLUXO 2: FAZ CONSULTA JOIN DE CURSO E CATEGORIA ====
	queries := db.New(dbConn)

	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}

	for _, course := range courses {
		fmt.Printf("Category: %s - Course ID: %s, Course Name: %s, Course Description: %s, Course Price: %f\n", course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
	}

}
