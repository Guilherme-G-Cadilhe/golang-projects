package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int
	Saldo  int
}

func main() {

	conta123 := Conta{Numero: 123, Saldo: 1000}

	// Serializa em JSON (Marshal = converte em JSON)
	res, err := json.Marshal(conta123)
	if err != nil {
		fmt.Println("error:", err)
	}
	// Converte os Bytes em String
	fmt.Println(string(res))

	// Cria um encoder (Encoder = codifica) e o envia para o Stdout (Stdout = saida padrao , terminal)
	encoder := json.NewEncoder(os.Stdout)
	// Define qual tipo de dados o encoder vai codificar
	err = encoder.Encode(conta123)
	// Versão resumida
	//	err = json.NewEncoder(os.Stdout).Encode(conta123)
	if err != nil {
		fmt.Println("error:", err)
	}

	jsonPuro := []byte(`{"numero": 567, "saldo": 2000}`)
	var contaX Conta

	// Deserializa em JSON (Unmarshal = converte de JSON para struct)
	// Faz a ligação com base no  tipo da Struct e nos dados do JSON
	// Se o JSON for inválido e for do tipo correto, ocorrera um erro
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Struct Conta: ", contaX, "\nSaldo:", contaX.Saldo, "\nNúmero:", contaX.Numero)

	type ContaComTags struct {
		Numero int `json:"n"` // Adiciona a propriedade "n" no JSON para o campo "Numero" quando houver conversão, então 'n' sera o nome do campo
		Saldo  int `json:"s"`
	}

	jsonPuroComTags := []byte(`{"n": 8910, "s": 3000}`)
	var contaY ContaComTags
	err = json.Unmarshal(jsonPuroComTags, &contaY)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Struct ContaComTags: ", contaY, "\ns:", contaY.Saldo, "\nn:", contaY.Numero)

}
