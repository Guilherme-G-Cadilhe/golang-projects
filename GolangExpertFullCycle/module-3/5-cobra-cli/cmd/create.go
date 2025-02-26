package cmd

import (
	"fmt"

	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-3/5-cobra-cli/internal/database"
	"github.com/spf13/cobra"
)

// newCreateCmd cria um novo comando "create" para o CLI.
// Recebe uma instância do repositório de categorias e configura o comando.
func newCreateCmd(categoryDb database.Category) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Cria Categoria",
		Long:  `Cria uma nova categoria no banco de dados.`,
		RunE:  runCreate(categoryDb), // Utiliza a função runCreate para isolar a lógica de execução.
	}
	return cmd
}

// runCreate retorna uma função compatível com RunEFunc que executa a lógica para criar uma categoria.
func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		// Recupera os valores das flags "name" e "description".
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		// Chama o método Create do repositório de categorias.
		category, err := categoryDb.Create(name, description)
		if err != nil {
			return err
		}
		// Exibe o ID da categoria criada.
		fmt.Printf("Category created with id: %s\n", category.ID)
		return nil
	}
}

func init() {
	// Cria um novo comando "create" utilizando a função newCreateCmd.
	createCmd := newCreateCmd(GetCategoryDB(GetDb()))
	// Adiciona o comando "create" como subcomando do comando "category" (presumivelmente definido em outro arquivo).
	categoryCmd.AddCommand(createCmd)
	// Define flags locais para o comando "create".
	createCmd.Flags().StringP("name", "n", "", "Nome da categoria")
	createCmd.Flags().StringP("description", "d", "", "Descrição da categoria")
	// Exige que ambas as flags "name" e "description" sejam fornecidas juntas.
	createCmd.MarkFlagsRequiredTogether("name", "description")
}
