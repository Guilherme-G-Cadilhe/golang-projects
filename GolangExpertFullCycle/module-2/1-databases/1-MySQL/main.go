package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Importa o driver MySQL para o pacote database/sql
	"github.com/google/uuid"           // Importa o pacote para gerar UUIDs
)

// Struct que representa um produto
type Product struct {
	ID    string  // ID do produto
	Name  string  // Nome do produto
	Price float64 // Preço do produto
}

// Função que cria um novo produto com um UUID gerado automaticamente
func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	// Abre uma conexão com o banco de dados MySQL
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err) // Encerra a execução do programa se ocorrer um erro
	}
	defer db.Close() // Fecha a conexão ao final da função main

	// Cria um novo produto
	product := NewProduct("Notebook VRAU", 1899.99)
	fmt.Println("Criando um produto", product)

	// Insere o produto no banco de dados
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	// Atualiza o preço do produto
	product.Price = 1200.99
	fmt.Println("Atualiza o produto", product)

	// Atualiza o produto no banco de dados
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	// Seleciona o primeiro produto do banco de dados
	seletecProduct, err := selectOneProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Consultando o produto", seletecProduct)

	// Seleciona todos os produtos do banco de dados
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Println("Consultando todos os produtos", p)
	}

	// Deleta o primeiro produto do banco de dados
	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deletando o produto", product)

}

// Função que insere um produto no banco de dados
func insertProduct(db *sql.DB, product *Product) error {
	// Prepara a instrução SQL para inserção
	// Cria uma declaração preparada que pode ser executada várias vezes com diferentes parâmetros. Isso melhora a performance e segurança, evitando injeção de SQL.
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close() // Fecha o statement ao final da função

	// Executa a instrução SQL com os valores do produto
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

// Função que atualiza um produto no banco de dados
func updateProduct(db *sql.DB, product *Product) error {
	// Prepara a instrução SQL para atualização
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close() // Fecha o statement ao final da função

	// Executa a instrução SQL com os valores atualizados do produto
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

// Função que seleciona um produto no banco de dados pelo ID
func selectOneProduct(db *sql.DB, id string) (*Product, error) {
	// Prepara a instrução SQL para seleção
	statement, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer statement.Close() // Fecha o statement ao final da função

	// Executa a instrução SQL e escaneia o resultado para a struct Product
	var p Product
	err = statement.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// Função que seleciona um produto com contexto para suporte a cancelamento e prazos
func selectOneProductWithContext(ctx context.Context, db *sql.DB, id string) (*Product, error) {
	statement, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer statement.Close() // Fecha o statement ao final da função

	var p Product
	err = statement.QueryRowContext(ctx, id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// Função que seleciona todos os produtos no banco de dados
func selectAllProducts(db *sql.DB) ([]Product, error) {
	// Executa a instrução SQL sem preparar um statement (não há parâmetros)
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Fecha o conjunto de resultados ao final da função

	var products []Product

	// Itera sobre todas as linhas retornadas
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

// Função que deleta um produto no banco de dados pelo ID
func deleteProduct(db *sql.DB, id string) error {
	// Prepara a instrução SQL para deleção
	statement, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close() // Fecha o statement ao final da função

	// Executa a instrução SQL para deletar o produto
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
