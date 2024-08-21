package main

import (
	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-2/2-packaging/3-goModReplace/math"
)

// Caso não esteja publicado e seja um modulo local que está tentando acesar, existe duas formas de utilizar:
// 1 Gambiarra local - 'go mod edit -replace github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-3/1-packaging/3-goModReplace/math=../math'
// rodar 'go mod tidy'

func main() {

	m := math.NewMath(1, 2)
	println(m.SumPrivate())
}
