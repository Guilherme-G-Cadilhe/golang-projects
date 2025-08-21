package main

import "fmt"

func add(a, b int) int {
	sum := a + b // sum é alocada na stack
	return sum   // liberada quando a função termina
}

type User struct {
	Name string
}

// A variavel user é um ponteiro para um objeto User
// O ponteiro user é retornado pela função createUser
// O objeto user precisa continuar existindo após o retorno da função, pois o ponteiro é usado fora do escopo da função, portanto ele precisa ser alocado na heap
func createUser(name string) *User {
	user := &User{Name: name} // user é alocada na heap porque retorna um ponteiro
	return user
}

// Armazenamento em Estruturas de Dados
// Se uma variavel local é armazenada em uma estrutura de dados que sobrevive ao esocpo da função, ela deve ser alocada an heap.

func storeInMap() map[string]*int {
	m := make(map[string]*int) // m é um mapa que pode sobreviver ao escopo da função
	i := 42                    // i é uma variavel local
	m["key"] = &i              // o ponteiro i é armazenado no mapa
	return m
	//IMPORTANTE: se m não fosse retornado, o ponteiro i seria liberado quando a funcao terminasse
}

func storeInMapWithNoReturn() {
	m := make(map[string]*int) // m não escapa para heap
	i := 42                    // i é uma variavel local
	m["key"] = &i              // o ponteiro i é armazenado no mapa
}

// Se uma variavel local é usada dentro de uma goroutine, ela deve ser alocada na heap, pois a goroutine pode continuar a execução após a função retornar.

func startWorker() {
	z := 10
	go func() {
		// Esta goroutine pode rodar muito depois de startWorker() ter retornado.
		// 'z' precisa continuar vivo na Heap.
		fmt.Println(z)
	}()
}

// função recursiva que poderia gerar um stack overflow
func recursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursive(n-1)
}

// dlv é um debuger para go
// go build -gcflags="-m" main.go
// go install github.com/go-delve/delve/cmd/dlv@latest
// go build -gcflags="all=-N -l" -o recursive // compila a funcao recursiva
// dlv debug // inicia o debuger
// break main.recursive // adiciona um breakpoint na funcao recursiva
// continue // executa a funcao recursiva
//step...step...step... // cada step executa uma linha da funcao
// stack // mostra a pilha de execucao

func main() {
	// fmt.Println(add(1, 2)) // fmt faz com que 1 e 2 escapassem para a heap
	println(add(1, 2)) // println não escapa para a heap

	var p *int
	{
		i := 42
		p = &i // i escapa para a heap porque p é uma variável de escopo maior
	}

	fmt.Println(*p)
	fmt.Println(recursive(10))

}
