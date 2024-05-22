package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Config map[string]interface{}

func loadConfig(jsonData string) (Config, error) {
	var config Config
	err := json.Unmarshal([]byte(jsonData), &config)
	return config, err
}

func processConfig(config Config) {

	for key, value := range config { // config ainda é um pointer de memoria
		switch v := value.(type) {
		case string:
			fmt.Printf("Config %s is a string: %s\n", key, v)
		case bool:
			fmt.Printf("Config %s is a bool: %v\n", key, v)
		case float64:
			fmt.Printf("Config %s is an number: %d\n", key, int(v))
		default:
			fmt.Printf("Config %s is of an unknown type\n", key)
		}
	}
}
func main() {
	var minhaVar interface{} = "Guilherme"

	println(minhaVar)
	println(minhaVar.(string)) // Output: Guilherme, consegue fazer a conversão e retorna o valor

	res, ok := minhaVar.(int)
	println(res, ok) // Output: 0 false, não consegue fazer a conversão por que é uma string
	// res2 := minhaVar.(int)
	// println(res2) // Lança error PANIC por que não tem o Catch do OK pra sinalizar erro

	jsonData := `{
    "appName": "MyApp",
    "debug": true,
    "maxConnections": 10
}`

	config, err := loadConfig(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	// Processar a configuração
	processConfig(config)

}
