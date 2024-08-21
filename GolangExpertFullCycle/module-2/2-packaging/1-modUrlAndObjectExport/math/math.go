package math

import "fmt"

// No Go, qualquer identificador (variável, struct, método) que começa com uma letra maiúscula é exportado, ou seja, pode ser acessado de outros pacotes. Identificadores com iniciais minúsculas são privados ao pacote.
var Public = "hello math externo"  // Exportavel
var private = "hello math interno" // Não exportavel

type MathPublic struct {
	A int
	B int
}

// Impede que as variaveis sejam alteraveis por fora, necessitando de uma função especifica que permita ela ser alterada
// O struct mathPrivate com campos não exportados (a, b, C) só pode ser manipulado dentro do pacote math. Para interagir com esses campos de fora do pacote, são fornecidos métodos como SumPrivate.
type mathPrivate struct {
	a int
	b int
	C int
}

func NewMath(a, b int) mathPrivate {
	return mathPrivate{a: a, b: b}
}

// No Go, métodos estão associados a structs, e só podem ser chamados em instâncias dessas structs. A função Sum é um método de MathPublic e só pode ser chamado em instâncias de MathPublic. Similarmente, SumPrivate é um método de mathPrivate, acessível apenas para instâncias de mathPrivate.
func (m MathPublic) Sum() int {
	fmt.Println(private)
	return m.A + m.B
}

func (m *mathPrivate) SumPrivate() int {
	return m.a + m.b
}
