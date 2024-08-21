package main

import (
	"fmt"

	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-3/1-packaging/1-modUrlAndObjectExpor/math"
)

func main() {
	// No Go, structs e métodos com iniciais maiúsculas são exportados, ou seja, podem ser acessados de fora do pacote. O struct MathPublic é inicializado e seu método Sum é chamado para retornar a soma dos valores A e B.
	m := math.MathPublic{A: 1, B: 2}

	fmt.Println(math.Public)

	fmt.Println(m.Sum())

	// O método NewMath cria e retorna uma instância do struct mathPrivate. A instância retornada pode então ser usada para acessar métodos exportados ou não exportados, dependendo da visibilidade. A função SumPrivate é um método que só pode ser chamado em uma instância de mathPrivate.
	m2 := math.NewMath(3, 4)
	m2.C = 5
	fmt.Println(m2.SumPrivate())
	fmt.Println(m2.C)

}
