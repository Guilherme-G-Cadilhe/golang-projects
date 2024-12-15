package tax

import "github.com/stretchr/testify/mock"

// TaxRepositoryMock é uma implementação mockada da interface Repository.
// Utiliza a biblioteca testify para facilitar a criação de mocks em testes.
type TaxRepositoryMock struct {
	mock.Mock
}

// SaveTax é a implementação mockada do método SaveTax da interface Repository.
func (m *TaxRepositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)
}
