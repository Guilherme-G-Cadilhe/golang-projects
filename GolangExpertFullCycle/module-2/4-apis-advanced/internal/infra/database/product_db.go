package database

import (
	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-2/4-apis-advanced/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product

	// err := p.DB.Where("id = ?", id).First(&product).Error
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error
	if sort != "" && sort != "desc" && sort != "asc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 { // Se os valores de page e limit forem diferentes de 0
		// Offset = (page - 1) * limit
		// Offset significa que o gorm irá pegar os dados a partir da posição 0 até a posição limit
		// Order = "created_at " + sort
		// Order significa que o gorm irá ordenar os dados pelo campo created_at em ordem ascendente ou descendente de acordo com o valor de sort adicionado no final da string
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}

	return products, err
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}
