package graph

import "github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-3/2-graphql/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Course
}
