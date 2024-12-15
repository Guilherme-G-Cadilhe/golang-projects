package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestCalculateTax verifica se a função CalculateTax retorna os valores esperados e trata erros corretamente.
func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)

	// Verifica se não houve erro ao calcular o imposto.
	assert.Nil(t, err)

	// Verifica se o valor do imposto está correto.
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0)

	// Verifica se a função retornou o erro esperado.
	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, tax)

	// Verifica se a mensagem de erro específica foi retornada.
	assert.EqualError(t, err, "amount must be greater than 0")
}

// TestCalculateTaxAndSave verifica se a função CalculateTaxAndSave está calculando e salvando o imposto corretamente.
func TestCalculateTaxAndSave(t *testing.T) {

	// Cria um mock do repositório.
	repository := &TaxRepositoryMock{}

	// Configura o mock para retornar nil quando o valor do imposto for 10.0.
	repository.On("SaveTax", 10.0).Return(nil)

	// Configura o mock para retornar nil uma única vez para 10.0.
	repository.On("SaveTax", 10.0).Return(nil).Once()

	// Configura o mock para retornar nil duas vezes para 10.0.
	repository.On("SaveTax", 10.0).Return(nil).Twice()

	// Configura o mock para retornar um erro quando o valor do imposto for 0.0.
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	// Configura o mock para retornar um erro para qualquer valor.
	repository.On("SaveTax", mock.Anything).Return(errors.New("error saving tax"))

	// Testa a função CalculateTaxAndSave com o valor 1000.0.
	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err) // Verifica se não houve erro.

	// Testa a função CalculateTaxAndSave com o valor 0.0.
	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax") // Verifica se houve erro.

	// Verifica se as expectativas configuradas no mock foram atendidas.
	repository.AssertExpectations(t)

	// Verifica se o método SaveTax foi chamado exatamente 2 vezes.
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}
