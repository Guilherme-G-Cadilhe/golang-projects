package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/valyala/fastjson"
)

// Definindo a estrutura Address com tags JSON para serialização e desserialização
type Address struct {
	Street string `json:"street"`
	Number int    `json:"number"`
}

func main() {
	// Inicializando o parser fastjson
	var parser fastjson.Parser

	// Definindo uma string JSON com dados fictícios
	jsonData := `{
	    "appName": "MyApp",
	    "debug": true,
	    "maxConnections": 10,
	    "address": {
	      "street": "Main Street",
	      "number": 123
	    },
	    "array": [1, 2, 3]
	}`

	// Parse a string JSON para obter um objeto fastjson.Value
	// value é do tipo *fastjson.Value, que representa o JSON analisado.
	value, err := parser.Parse(jsonData)
	if err != nil {
		panic(err)
	}
	fmt.Println(value)

	// Obtendo uma string do JSON
	fmt.Printf("appName=%s\n", string(value.GetStringBytes("appName")))
	// Obtendo um boolean do JSON
	fmt.Printf("debug=%t\n", value.GetBool("debug"))
	// Obtendo um inteiro do JSON
	fmt.Printf("maxConnections=%d\n", value.GetInt("maxConnections"))
	// Obtendo um campo aninhado (número do endereço)
	fmt.Printf("number=%d\n", value.GetInt("address.number"))

	// Obtendo e iterando sobre um array do JSON
	a := value.GetArray("array")
	for i, v := range a {
		fmt.Printf("array[%d]: %s\n", i, v)
	}

	// Obtendo um objeto do JSON (método 1)
	address := value.GetObject("address")
	fmt.Printf("address street: %s\n", address.Get("street"))
	fmt.Printf("address number: %s\n", address.Get("number"))

	// Obtendo um objeto do JSON (método 2) e desserializando em uma struct Go
	// Get("address").String(): Obtém o objeto JSON como string.
	address2JSON := value.Get("address").String()
	var address2STRUCT Address
	if err := json.Unmarshal([]byte(address2JSON), &address2STRUCT); err != nil {
		panic(err)
	}
	fmt.Println("address2 street:", address2STRUCT.Street)
	fmt.Println("address2 number:", address2STRUCT.Number)

	// Inicializando o scanner fastjson com uma string JSON
	// fastjson.Scanner é usado para escanear tokens individuais em uma string JSON.
	var sc fastjson.Scanner
	sc.Init(`   {"foo":  "bar"  }[  ]
		12345"xyz" true false null    `)

	// Iterando sobre os tokens na string JSON
	// sc.Next: Move para o próximo token.
	for sc.Next() {
		// sc.Value: Obtém o valor do token atual.
		fmt.Printf("%s\n", sc.Value())
	}
	// Verificando se houve algum erro durante a análise dos tokens
	if err := sc.Error(); err != nil {
		log.Fatalf("unexpected error: %s", err)
	}
}
