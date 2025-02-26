package cmd

import (
	"database/sql"
	"os"

	// Importa o módulo de acesso ao banco e o pacote gerado pelo Cobra (além do driver SQLite)
	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-3/5-cobra-cli/internal/database"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// RunEFunc é um tipo customizado que representa uma função de execução que pode retornar erro.
// Isso é similar a definir uma função assíncrona que retorna uma Promise em JavaScript.
type RunEFunc func(cmd *cobra.Command, args []string) error

// GetDb encapsula a abertura da conexão com o banco de dados SQLite.
// Funciona como um módulo de conexão no Node.js (ex.: usando o sqlite3 ou Sequelize).
func GetDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	return db
}

// GetCategoryDB retorna a instância do repositório (ou model) de categoria, pronta para uso.
func GetCategoryDB(db *sql.DB) database.Category {
	return *database.NewCategory(db)
}

// rootCmd é o comando base da aplicação CLI.
// Em frameworks JavaScript como Commander, seria equivalente ao comando principal da CLI.
var rootCmd = &cobra.Command{
	Use:   "5-cobra-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Se nenhum subcomando for especificado, o comando principal pode exibir a ajuda.
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute invoca o comando raiz e trata eventuais erros, similar a chamar "program.parse(process.argv)" no Commander.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init é executado automaticamente para configurar flags e outros parâmetros globais.
// Aqui, adiciona uma flag local (não persistente) que pode ser usada para controle extra.
func init() {
	// Exemplo de flag local; flags persistentes afetariam todos os subcomandos.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
