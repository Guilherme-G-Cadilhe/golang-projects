// > go test -fuzz=. (Roda os testes de Fuzzing)
// > go test -fuzz=. -fuzztime=5s (Tempo limite fazendo Fuzzing)

package tax

import "testing"

// Função de teste de Fuzzing que utiliza a estrutura `testing.F`
// O Fuzzing é uma técnica de teste que fornece entradas aleatórias ou inesperadas para tentar encontrar bugs em uma função.
// Nesse caso, estamos testando a função `CalculateTax`.

func FuzzCalculateTax(f *testing.F) {

	// Seed é um slice de floats que contém valores iniciais usados como ponto de partida para o Fuzzing.
	// Incluímos valores negativos, zero, e valores típicos que a função `CalculateTax` deve lidar.
	seed := []float64{-1, -2, -2.5, 0.0, 500.0, 1000.0, 1500.0}

	// Adicionamos cada valor da seed ao fuzzer para que ele os utilize como base para gerar novos valores.
	for _, amount := range seed {
		f.Add(amount) // Adiciona o valor atual do seed ao fuzzer.
	}

	// A função Fuzz é onde o fuzzer começa a gerar entradas aleatórias com base nos valores iniciais fornecidos.
	// O método recebe dois parâmetros: `t *testing.T` que representa o contexto do teste e `amount float64` que é a entrada gerada.
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		if amount <= 0 && result != 0 {
			t.Errorf("Received %f, but expected 0", result)
		}

		if amount > 20000 && result != 20 {
			t.Errorf("Received %f, but expected 20", result)
		}
	})

	// Passa multiplos parametros
	// for _, amount := range seed {
	// 	f.Add(amount, 10) // A e B
	// }
	// f.Fuzz(func(t *testing.T, a float64, b float64) {
	// 	result := CalculateTax(a, b)
	// })
}
