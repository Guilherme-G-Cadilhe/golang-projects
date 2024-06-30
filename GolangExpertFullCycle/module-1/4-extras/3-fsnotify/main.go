package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

type DBConfig struct {
	DB       string `json:"db"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var config DBConfig

func main() {

	// Cria um novo fsnotify watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	MarshalConfig("config.json")

	// Cria um canal para receber notificações
	done := make(chan bool)
	go func() {
		for {
			select {
			//  Escuta eventos no watcher
			case event, ok := <-watcher.Events:
				if !ok { // Não achou eventos
					// Retorna para o For
					return
				}
				fmt.Println("event:", event)
				//  Verifica se o evento é de escrita no arquivo monitorado.
				if event.Op&fsnotify.Write == fsnotify.Write {
					MarshalConfig("config.json")
					fmt.Println("Config depois >", config)
				}

			// Escuta erros no watcher.
			case err, ok := <-watcher.Errors:
				if !ok { // Não achou erros
					// Retorna para o For
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()

	// Cria um novo watcher
	err = watcher.Add("config.json")
	if err != nil {
		panic(err)
	}
	<-done

}

// MarshalConfig le o arquivo JSON de configuração e converte o JSON para uma struct
func MarshalConfig(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
}
