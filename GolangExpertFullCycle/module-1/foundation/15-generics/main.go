package main

func SomaInteiro(m map[string]int) int {
	var soma int

	// Chave, Valor := percorrer cada propriedade do map (Estilo Objeto em JSON)
	for _, v := range m {
		soma += v
	}

	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64

	for _, v := range m {
		soma += v
	}

	return soma
}

// Usando Generics
func Soma[T int | float64](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}
	return soma
}

type MyNumber int

// o ~ indicar tipos subjacentes, permitindo que a interface aceite tipos que têm um tipo subjacente específico.
type Number interface {
	~int | float64
}

func SomaComInterface[T Number](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}
	return soma
}

// Comparable permite comparar valores
func Compara[T comparable](a, b T) bool {
	return a == b
}

func main() {

	numeros := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	numerosFloat := map[string]float64{
		"a": 1.5,
		"b": 2.5,
		"c": 3.5,
		"d": 4.5,
	}
	numeros2 := map[string]MyNumber{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	println(SomaInteiro(numeros))
	println(SomaFloat(numerosFloat))

	println(Soma(numeros))
	println(Soma(numerosFloat))

	println(SomaComInterface(numeros))
	println(SomaComInterface(numeros2))

	println(Compara(2, 2))

}
