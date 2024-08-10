package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []GormProduct
}
type GormProduct struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category     `gorm:"foreignKey:CategoryID"`
	SerialNumber SerialNumber `gorm:"foreignKey:ProductID;references:ID"`
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&GormProduct{}, &Category{}, &SerialNumber{})

	// create category
	// category := Category{Name: "Cozinha"}
	// db.Create(&category)

	// // create product
	// db.Create(&GormProduct{
	// 	Name:       "Panela",
	// 	Price:      300.00,
	// 	CategoryID: 1,
	// })

	// db.Create(&SerialNumber{
	// 	ProductID: 1,
	// 	Number:    "123456789",
	// })

	var categories []Category
	// err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	// Utilizar Preload com categorias aninhadas já puxa o ambos, pai e filho
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	// fmt.Println(categories)
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println(" - [Product]: ", product.Name, "- [Serial Number]:"+product.SerialNumber.Number)
		}
	}
}
