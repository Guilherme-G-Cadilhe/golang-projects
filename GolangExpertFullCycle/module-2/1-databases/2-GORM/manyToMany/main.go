package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []GormProduct `gorm:"many2many:products_categories;"`
	// Products []GormProduct `gorm:"many2many:products_categories;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
type GormProduct struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	// Categories []Category `gorm:"many2many:products_categories;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&GormProduct{}, &Category{})

	// create category
	// category := Category{Name: "Cozinha"}
	// db.Create(&category)

	// category2 := Category{Name: "Eletronicos"}
	// db.Create(&category2)

	// // create product
	// db.Create(&GormProduct{
	// 	Name:       "Airfryer",
	// 	Price:      300.00,
	// 	Categories: []Category{category, category2},
	// })

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	// fmt.Println(categories)
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println(" - [Product]: ", product.Name)
		}
	}
}
