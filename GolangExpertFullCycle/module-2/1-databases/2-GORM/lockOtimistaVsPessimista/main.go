package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

/*
- Lock otimista:
 No lock otimista, não há bloqueio explícito. Em vez disso, utiliza-se uma versão dos dados para detectar se ocorreu uma atualização concorrente. Se a versão do registro mudou desde o momento em que a transação começou, a transação falha.

name, email, versão
W     W@W       1    - Inicial
W     W@W       1    - Finalizou a transação
W     W@W       2    - Versão incrementada
W     W@W       1    - Finalizou a transação [Não pode]

- Lock pessimista:
No lock pessimista, a linha ou tabela é bloqueada pelo banco de dados para garantir que nenhuma outra transação possa modificar os dados até que o bloqueio seja liberado. Isso é útil para evitar condições de corrida quando se está atualizando registros.

Exemplo em Mysql:
select * from products where id = 1 FOR UPDATE [ Bloca a linha]
update products set name = 'X' where id = 1
*/

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&GormProduct{}, &Category{}) // Cria tabelas baseadas nas estruturas GormProduct e Category

	// Iniciando uma transação para demonstração de lock pessimista
	tx := db.Begin() // Inicia uma transação

	var c Category
	// Aplicando lock pessimista na transação, e pegando o primeiro registro
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Casas"    // Modificando o nome da categoria
	tx.Debug().Save(&c) // Salvando a categoria modificada
	tx.Commit()         // Finalizando a transação e liberando o lock pessimista

}
