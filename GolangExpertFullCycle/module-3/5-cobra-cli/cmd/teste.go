package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// testeCmd representa o comando "teste" que pode ser chamado via CLI.
var testeCmd = &cobra.Command{
	Use:   "teste",
	Short: "A brief description of your command",
	Long:  ``,
	// Run é executado quando o comando é chamado; aqui, demonstra a leitura de flags e a execução de lógica.
	Run: func(cmd *cobra.Command, args []string) {
		// Recupera a flag "comando" e decide o comportamento (semelhante a process.argv ou process.env em JS).
		comando, _ := cmd.Flags().GetString("comando")
		if comando == "ping" {
			fmt.Println("Ping")
		} else if comando == "pong" {
			fmt.Println("Pong")
		} else {
			fmt.Println("Escolha ping ou pong")
		}

		// Exibe o valor da flag persistente "name", armazenado na variável global categoryTeste.
		fmt.Println("Category called with name:", categoryTeste)

		// Exibe o valor das flags "exists" e "id".
		exists, _ := cmd.Flags().GetBool("exists")
		fmt.Println("Category called with exists:", exists)

		id, _ := cmd.Flags().GetInt16("id")
		fmt.Println("Category called with id:", id)
	},
	// PreRun é executado antes do Run, ideal para pré-processamento; similar a "before" em Mocha.
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chamado antes do run")
	},
	// PostRun é executado após o Run, permitindo ações de limpeza ou logging.
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chamado depois do run")
	},
	// RunE permite retornar um erro; útil para fluxos onde você precisa lidar com falhas.
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("error")
	},
}

// Variável global para armazenar o valor da flag persistente "name".
var categoryTeste string

func init() {
	// Adiciona o comando "teste" ao comando raiz.
	rootCmd.AddCommand(testeCmd)
	// Define a flag "comando" com atalho "-c".
	testeCmd.Flags().StringP("comando", "c", "", "Escolha ping ou pong")
	// Torna obrigatória a flag "comando".
	testeCmd.MarkFlagRequired("comando")

	// Define uma flag persistente que fica disponível em todos os subcomandos de "teste"
	// e armazena seu valor na variável categoryTeste.
	testeCmd.PersistentFlags().StringVarP(&categoryTeste, "name", "n", "Y", "Nome da categoria")

	// Define a flag "exists", que é booleana.
	testeCmd.PersistentFlags().BoolP("exists", "e", false, "Checa se categoria existe")
	// Define a flag "id" como um inteiro de 16 bits.
	testeCmd.PersistentFlags().Int16P("id", "i", 0, "ID da categoria")
}
