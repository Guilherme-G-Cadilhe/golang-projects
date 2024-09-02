package tax

import "testing"

// > go test -fuzz=. (Roda os testes de Fuzzing)
// > go test -fuzz=. -fuzztime=5s (Tempo limite fazendo Fuzzing)

func FuzzCalculateTax(f *testing.F) {

	seed := []float64{-1, -2, -2.5, 0.0, 500.0, 1000.0, 1500.0}

	// Passa multiplos parametros
	// for _, amount := range seed {
	// 	f.Add(amount, 10) // A e B
	// }
	// f.Fuzz(func(t *testing.T, a float64, b float64) {
	// 	result := CalculateTax(a, b)
	// })

	// Passa um unico metodo
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f, but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Received %f, but expected 20", result)
		}
	})

}
