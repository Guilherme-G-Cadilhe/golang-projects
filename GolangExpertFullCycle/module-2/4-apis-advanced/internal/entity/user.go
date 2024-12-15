package entity

import (
	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-2/4-apis-advanced/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

// VO - Value Object
type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"` //Nunca exibe o password para o usuário final
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
