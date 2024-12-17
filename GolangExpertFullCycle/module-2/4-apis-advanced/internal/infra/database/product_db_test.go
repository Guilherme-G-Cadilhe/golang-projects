package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-2/4-apis-advanced/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)
	product, err := entity.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, productDB.Create(product))
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	for i := 1; i <= 24; i++ {
		// Utiliza o fmt.Sprintf para criar um nome de produto dinâmico
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.Nil(t, err)
		assert.NotNil(t, product)
		assert.Nil(t, db.Create(product).Error)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 10, len(products))
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	// fmt.Println(products)
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 4)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 24", products[3].Name)

}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)
	product, err := entity.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, productDB.Create(product))
	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.NotNil(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)
	product, err := entity.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, productDB.Create(product))
	product.Name = "Product 2"
	product.Price = 20
	assert.Nil(t, productDB.Update(product))
	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.NotNil(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)
	product, err := entity.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, productDB.Create(product))

	assert.Nil(t, productDB.Delete(product.ID.String()))

	_, err = productDB.FindByID(product.ID.String())
	assert.NotNil(t, err)
}
