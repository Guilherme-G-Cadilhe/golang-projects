package main

import (
	"fmt"

	"gorm.io/driver/mysql" // Importa o driver MySQL para GORM
	"gorm.io/gorm"         // Importa o pacote GORM
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

// GormProduct define a estrutura do produto para o GORM, mapeada para a tabela "gorm_products"
type GormProduct struct {
	ID         int     `gorm:"primaryKey"` // Define a chave primária
	Name       string  // Nome do produto
	Price      float64 // Preço do produto
	gorm.Model         // Definição de campos adicionais (CreatedAt, UpdatedAt, DeletedAt)
}

func main() {
	// Definindo o DSN (Data Source Name) para conectar ao banco de dados MySQL
	// Exemplo: username:password@protocol(address)/dbname?param1=value1&param2=value2
	/*
		Parâmetros Opcionais Comuns
		charset: Define o conjunto de caracteres. Por exemplo, utf8mb4 suporta uma gama completa de caracteres Unicode.
		parseTime: Quando True, instrui o driver a analisar os valores date e datetime como tipos time.Time do Go.
		loc: Define o fuso horário. Pode ser Local ou um valor específico de timezone.
		timeout: Define o tempo limite para a tentativa de conexão, por exemplo, timeout=30s.
		readTimeout: Define o tempo limite para operações de leitura, por exemplo, readTimeout=30s.
		writeTimeout: Define o tempo limite para operações de escrita, por exemplo, writeTimeout=30s.
		tls: Configura a conexão TLS/SSL, por exemplo, tls=true.
	*/
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	// Abre uma conexão com o banco de dados MySQL usando GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err) // Encerra a execução do programa se ocorrer um erro na conexão
	}

	// AutoMigrate cria a tabela "gorm_products" se ela ainda não existir
	// O método AutoMigrate cria a tabela com base na estrutura GormProduct
	db.AutoMigrate(&GormProduct{})

	// Criação de produtos
	createProducts(db)

	// Leitura de produtos
	readProducts(db)

	// Atualização de produtos
	updateProduct(db)

	// Exclusão de produtos
	deleteProduct(db)
}

// Função para criar produtos
func createProducts(db *gorm.DB) {
	// Cria um novo produto na tabela "gorm_products"
	db.Create(&GormProduct{
		Name:  "Mouse", // Nome do produto
		Price: 200.99,  // Preço do produto
	})

	// Cria múltiplos produtos em batch (lote)
	products := []GormProduct{
		{Name: "Mouse VRAU", Price: 2000.99},
		{Name: "Monitor VRAU", Price: 889.99},
		{Name: "Notebook BLA", Price: 1000.99},
	}
	// Insere todos os produtos da slice products na tabela "gorm_products"
	db.Create(&products)
}

// Função para ler produtos
func readProducts(db *gorm.DB) {
	// Seleciona um registro na tabela "gorm_products" pelo ID
	var product GormProduct
	db.First(&product, 1) // Busca pelo ID 1
	fmt.Println("Produto com ID 1:", product)

	// Seleciona um registro na tabela "gorm_products" pelo nome
	var productByName GormProduct
	db.First(&productByName, "Name = ?", "Notebook TAR")
	fmt.Println("Produto com nome 'Notebook TAR':", productByName)

	// Seleciona todos os registros na tabela "gorm_products"
	var products []GormProduct
	db.Find(&products)
	fmt.Println("Todos os produtos:")
	for _, product := range products {
		fmt.Println(product)
	}

	// Seleciona os registros da tabela "gorm_products" com limite e offset
	// Pula 2 registros e seleciona os 2 próximos
	var limitedProducts []GormProduct
	db.Limit(2).Offset(2).Find(&limitedProducts)
	fmt.Println("Produtos com limite e offset:")
	for _, product := range limitedProducts {
		fmt.Println(product)
	}

	// Usando WHERE
	var productsWhere []GormProduct
	db.Where("price > ?", 1900).Find(&productsWhere)
	fmt.Println("Produtos com preço maior que 1900:")
	for _, product := range productsWhere {
		fmt.Println(product)
	}

	// Usando LIKE
	var productsLike []GormProduct
	db.Where("name LIKE ?", "%book%").Find(&productsLike)
	fmt.Println("Produtos com 'book' no nome:")
	for _, product := range productsLike {
		fmt.Println(product)
	}
}

// Função para atualizar produtos
func updateProduct(db *gorm.DB) {
	// Atualiza um registro na tabela "gorm_products"
	var p GormProduct
	db.First(&p, 2) // Busca pelo ID 2
	p.Name = "Mouse gamer"
	db.Save(&p)

	// Consulta o registro alterado
	var p2 GormProduct
	db.First(&p2, 2) // Busca pelo ID 2 novamente para verificar a atualização
	fmt.Println("Produto atualizado:", p2)
}

// Função para deletar produtos
func deleteProduct(db *gorm.DB) {
	// Deleta um registro na tabela "gorm_products"
	var p2 GormProduct
	db.First(&p2, 2) // Busca pelo ID 2
	db.Delete(&p2)
	fmt.Println("Produto deletado:", p2)
}
