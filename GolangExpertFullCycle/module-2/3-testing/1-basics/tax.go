package tax

import "time"

// A função `CalculateTax` recebe um valor (`amount`) e calcula o imposto baseado nesse valor.
// Se o valor for 0, retorna 0.0, indicando que não há imposto.
// Se o valor for maior ou igual a 1000, retorna 10.0 como imposto fixo.
// Caso contrário, retorna 5.0.
func CalculateTax(amount float64) float64 {
	if amount == 0 {
		return 0.0
	}
	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}

// A função `CalculateTax2` é semelhante à `CalculateTax`, mas inclui um `time.Sleep` de 1 milissegundo.
// Isso pode ser útil para simular um processamento mais longo em testes de benchmark.
func CalculateTax2(amount float64) float64 {
	time.Sleep(time.Millisecond)
	if amount == 0 {
		return 0.0
	}
	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}
