package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCalculateTax verifica se a função CalculateTax retorna os valores esperados e trata erros corretamente.
func TestCalculateTax(t *testing.T) {
	testCases := []struct {
		name          string
		input         float64
		expectedTax   float64
		expectedError string
	}{
		{"Valid amount 1000", 1000.0, 10.0, ""},
		{"Valid amount 500", 500.0, 5.0, ""},
		{"Valid amount 20000", 20000.0, 20.0, ""},
		{"Invalid amount 0", 0.0, 0.0, "amount must be greater than 0"},
		{"Invalid amount -100", -100.0, 0.0, "amount must be greater than 0"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tax, err := CalculateTax(tc.input)

			if tc.expectedError != "" {
				assert.EqualError(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectedTax, tax)
		})
	}
}

// TestCalculateTaxAndSave verifica se a função CalculateTaxAndSave está calculando e salvando o imposto corretamente.
func TestCalculateTaxAndSave(t *testing.T) {
	testCases := []struct {
		name          string
		input         float64
		expectedTax   float64
		mockBehavior  func(*TaxRepositoryMock)
		expectedError string
	}{
		{
			name:        "Valid amount 1000",
			input:       1000.0,
			expectedTax: 10.0,
			mockBehavior: func(repo *TaxRepositoryMock) {
				repo.On("SaveTax", 10.0).Return(nil).Once()
			},
			expectedError: "",
		},
		{
			name:        "Invalid amount 0",
			input:       0.0,
			expectedTax: 0.0,
			mockBehavior: func(repo *TaxRepositoryMock) {
				repo.On("SaveTax", 0.0).Return(errors.New("error saving tax")).Once()
			},
			expectedError: "error saving tax",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repository := &TaxRepositoryMock{}
			tc.mockBehavior(repository)

			err := CalculateTaxAndSave(tc.input, repository)

			if tc.expectedError != "" {
				assert.EqualError(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}

			repository.AssertExpectations(t)
		})
	}
}

// Exemplo de teste de tabela para calculateTax
func TestCalculateTaxTable(t *testing.T) {
	testCases := []struct {
		amount float64
		want   float64
	}{
		{0, 0},
		{500, 5},
		{1000, 10},
		{15000, 10},
		{20000, 20},
		{25000, 20},
	}

	for _, tc := range testCases {
		got := calculateTax(tc.amount)
		assert.Equal(t, tc.want, got, "calculateTax(%f) = %f; want %f", tc.amount, got, tc.want)
	}
}

// Exemplo de teste de benchmark
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculateTax(1000)
	}
}
