package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	var cmd *exec.Cmd

	// runtime.GOOS detecta o sistema operacional
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "dir", "c:\\")
	} else {
		cmd = exec.Command("ls", "-l", "/")
	}

	// CombinedOutput() executa o comando e captura a saída padrão e de erro juntas
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao executar comando:", err)
		return
	}

	// Imprime a saída do comando executado
	fmt.Println("Saída do processo filho:")
	fmt.Println(string(output))
}
