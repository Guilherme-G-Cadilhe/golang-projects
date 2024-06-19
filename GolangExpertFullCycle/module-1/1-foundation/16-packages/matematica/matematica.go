package matematica

func Soma[T int | float64](a T, b T) T {
	return a + b
}

var A int = 10

// Maisculo = export, até mesmo para a propriedade "Marca", se ela não for, ela não fica visível de fora
type Carro struct {
	Marca string
}

// Função não está visível de fora
func (c Carro) ligar() string {
	return "Carro ligado"
}
