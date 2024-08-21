package main

import (
	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-3/1-packaging/4-goWorkspaces/math"
	"github.com/google/uuid"
)

// Quando você está trabalhando com múltiplos módulos Go localmente, você pode utilizar Workspaces para gerenciar
// e desenvolver esses módulos simultaneamente. Isso permite que você trabalhe com dependências locais sem a necessidade
// de publicar cada módulo em um repositório remoto.

// 1. Criar um workspace - Na pasta parente, rodar 'go work init ./math ./sistema'
// 2. Quando utilizando workspaces, o comando `go mod tidy` pode tentar buscar dependências remotamente, o que pode não ser desejado ao trabalhar com módulos locais. Nesse caso, use `go get` para instalar ou atualizar dependências.
//Alternativamente, `go mod tidy -e` pode ser usado para ignorar erros relacionados à importação de dependências locais.

// Resumo:
//   - Utilize `go work init` para criar um workspace que agrupe múltiplos módulos locais.
//   - Prefira `go get` em vez de `go mod tidy` ao gerenciar dependências dentro de um workspace para evitar conflitos com dependências locais.
func main() {

	m := math.NewMath(1, 2)
	println(m.SumPrivate())

	println(uuid.New().String())
}
