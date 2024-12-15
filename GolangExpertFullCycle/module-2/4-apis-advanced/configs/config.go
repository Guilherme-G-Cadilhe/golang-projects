package configs

import (
	"github.com/go-chi/jwtauth" // Biblioteca para criar e validar tokens JWT.
	"github.com/spf13/viper"    // Biblioteca para gerenciar configurações de arquivos e variáveis de ambiente.
)

// Define a estrutura para armazenar as configurações. Cada campo está mapeado para uma variável de ambiente específica.
type conf struct {
	DBDriver      string           `mapstructure:"DB_DRIVER"`       // Driver do banco de dados (ex: postgres, mysql).
	DBHost        string           `mapstructure:"DB_HOST"`         // Endereço do host do banco de dados.
	DBPort        string           `mapstructure:"DB_PORT"`         // Porta do banco de dados.
	DBName        string           `mapstructure:"DB_NAME"`         // Nome do banco de dados.
	DBUser        string           `mapstructure:"DB_USER"`         // Usuário do banco de dados.
	DBPassword    string           `mapstructure:"DB_PASSWORD"`     // Senha do banco de dados.
	WebServerPort string           `mapstructure:"WEB_SERVER_PORT"` // Porta do servidor web.
	JWTSecret     string           `mapstructure:"JWT_SECRET"`      // Chave secreta usada para assinar tokens JWT.
	JWTExpiresIn  int              `mapstructure:"JWT_EXPIRESIN"`   // Tempo de expiração dos tokens JWT (em segundos).
	TokenAuth     *jwtauth.JWTAuth // Objeto gerador de tokens JWT.
}

// Função que carrega as configurações a partir do arquivo .env
func LoadConfig(path string) (*conf, error) {
	var config *conf

	// Define o nome do arquivo de configuração (sem extensão).
	viper.SetConfigName("app_config")

	// Define o tipo do arquivo de configuração como '.env'.
	viper.SetConfigType("env")

	// Adiciona o caminho onde o arquivo de configuração está localizado.
	viper.AddConfigPath(path)

	// Força o carregamento do arquivo específico '.env'.
	viper.SetConfigFile(".env")

	// Habilita a leitura automática de variáveis de ambiente.
	viper.AutomaticEnv()

	// Tenta ler o arquivo de configuração. Caso falhe, o programa encerra.
	err := viper.ReadInConfig()
	if err != nil {
		// `panic` interrompe o programa se o arquivo não for encontrado ou tiver problemas de leitura.
		panic(err)
	}

	// Deserializa os valores do arquivo de configuração para a struct 'conf'.
	err = viper.Unmarshal(&config)
	if err != nil {
		// Retorna um erro se a deserialização falhar.
		return nil, err
	}

	// Inicializa o gerador de tokens JWT com a chave secreta especificada nas configurações.
	config.TokenAuth = jwtauth.New("HS256", []byte(config.JWTSecret), nil)

	// Retorna a configuração carregada e nenhum erro.
	return config, nil
}
