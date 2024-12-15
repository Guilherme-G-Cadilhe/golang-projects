package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Jogno", "jogon@go.dev", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Jogno", user.Name)
	assert.Equal(t, "jogon@go.dev", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Jogno", "jogon@go.dev", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("654321"))
	assert.NotEqual(t, "123456", user.Password)
}
