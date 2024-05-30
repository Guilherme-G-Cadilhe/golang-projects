package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("test.txt")
	if err != nil {
		println("error creating file")
		panic(err)
	}

	tamanho, err := f.Write([]byte("Hello, World! Its good to be here, have a nice day")) // Se for um slice de bytes
	// tamanho, err := f.WriteString("Hello, World!") // Se souber que vai ser String
	if err != nil {
		println("error writing string")
		panic(err)
	}
	fmt.Printf("wrote %d bytes\n", tamanho)
	f.Close() // Função que fechará o arquivo e executará as modificações

	// Leitura de interface do arquivo
	arquivoInterface, err := os.Open("test.txt") // >&{0xc00006ec88} (Pointer) e acesso a diferentes métodos do arquivo
	if err != nil {
		println("error opening file")
		panic(err)
	}
	fmt.Println("Arquivo nome:", arquivoInterface.Name())
	arquivoInterface.Close()

	// Leitura de conteudo do arquivo
	arquivoLeitura, err := os.ReadFile("test.txt") // >[72 101 108 108 111 44 32 87 111 114 108 100] (ByteArray)
	if err != nil {
		println("error reading file")
		panic(err)
	}
	fmt.Println("Arquivo conteudo:", string(arquivoLeitura)) // ByteArray para String

	// Leitura de pouco em pouco abrindo o arquivo
	arquivoStream, err := os.Open("test.txt")
	if err != nil {
		println("error read streaming file")
		panic(err)
	}
	// Criação da interface de reader por stream de dados
	reader := bufio.NewReader(arquivoStream)
	// Limite de 10 bytes a serem lidos a cada vez, cada iteração substitui os bytes antigos com novos 10
	smallBuffer := make([]byte, 10)
	for {
		// Escreve 10 bytes na variável smallBuffer a cada iteração, ate chegar no EOF (EOF = end of file)
		numberOfBytes, err := reader.Read(smallBuffer)
		if err != nil {
			break
		}
		// Imprime os 10 bytes atuais no buffer
		fmt.Println("Buffer: ", string(smallBuffer[:numberOfBytes]))
	}
	arquivoStream.Close()

	// Deleta o arquivo
	err = os.Remove("test.txt")
	if err != nil {
		println("error removing file")
		panic(err)
	}

}
