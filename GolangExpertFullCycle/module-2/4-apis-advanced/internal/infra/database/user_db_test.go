package database

import (
	"testing"

	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-2/4-apis-advanced/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	// Cria SQLite na memoria rapidamente pra testes
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{}) //Cria automaticamente as tabelas
	user, _ := entity.NewUser("Jogno", "jogon@go.dev", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)
	assert.NotEmpty(t, user.ID)

	var userFound entity.User                           //Cria um objeto vazio
	err = db.First(&userFound, "id = ?", user.ID).Error // Busca o usuário criado e armazena na variável userFound
	// t.Log(userFound)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Jogno", "jogon@go.dev", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)
	assert.NotEmpty(t, user.ID)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)

}
