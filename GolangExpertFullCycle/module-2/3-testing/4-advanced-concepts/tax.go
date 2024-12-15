package tax

import "errors"

// Repository define uma interface para salvar um valor de imposto.
// Essa abstração permite a injeção de diferentes implementações, facilitando testes.
type Repository interface {
	SaveTax(amount float64) error
}

// CalculateTaxAndSave calcula o imposto e tenta salvá-lo usando o repositório fornecido.
// Retorna um erro caso a operação de salvar falhe.
func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := calculateTax(amount)
	return repository.SaveTax(tax)
}

// calculateTax é uma função interna que calcula o valor do imposto baseado em regras específicas.
// Esta função substitui CalculateTax2, tornando-a privada ao pacote.
func calculateTax(amount float64) float64 {
	switch {
	case amount <= 0:
		return 0.0
	case amount >= 1000 && amount < 20000:
		return 10.0
	case amount >= 20000:
		return 20.0
	default:
		return 5.0
	}
}

// CalculateTax calcula o imposto e retorna um erro caso o valor fornecido seja inválido.
// Esta função é mantida para compatibilidade com código existente, mas considera-se o uso de calculateTax internamente.
func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0.0, errors.New("amount must be greater than 0")
	}
	return calculateTax(amount), nil
}
