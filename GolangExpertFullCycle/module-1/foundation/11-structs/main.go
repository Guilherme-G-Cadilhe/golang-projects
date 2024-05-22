package main

import "fmt"

type Teste struct {
	Teste1 string
	Teste2 int
}
type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}
type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
	Teste
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func main() {

	guilherme := Cliente{
		Nome:  "Gui",
		Idade: 30,
		Ativo: true,
	}

	guilherme.Teste1 = "Teste1"
	guilherme.Teste2 = 10
	guilherme.Address.Cidade = "São Paulo"
	guilherme.Address.Numero = 1000
	fmt.Println(guilherme)

	guilherme.Desativar()
	fmt.Println(guilherme)

}
