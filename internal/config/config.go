package config

import (
	"fmt"
	"os"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Server ServerConfig
	Database DatabaseConfig
	JWT JWTConfig
}

type ServerConfig struct {
	Port string
	Env string
}

type DatabaseConfig struct {
	Host string
	Port string
	User string
	Password string
	DBName string
	SSLMode string
}

type JWTConfig struct {
	Secret string
	ExpirationTime string
}

//Load carrega as configurações do arquivo .env
func Load() (*Config, error) {
	//Carrega o arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado.")
	}

	//Carrega as configurações do arquivo .env
	config := &Config{
        Server: ServerConfig{
            Port: getEnv("PORT", "8080"),
            Env: getEnv("ENV", "development"),
        },
        Database: DatabaseConfig{
            Host: getEnv("DB_HOST", "localhost"),
            Port: getEnv("DB_PORT", "5432"),
            User: getEnv("DB_USER", "postgres"),
            Password: getEnv("DB_PASSWORD", "postgres"),
            DBName: getEnv("DB_NAME", "smarttodo_db"),
            SSLMode: getEnv("DB_SSLMODE", "disable"),
        },
        JWT: JWTConfig{
            Secret: getEnv("JWT_SECRET", ""),
            ExpirationTime: getEnv("JWT_EXPIRATION", "24h"),
        },
    }


	// Validações do JWT
	if config.JWT.Secret == "" {
		return nil, fmt.Errorf("JWT_SECRET não definido no arquivo .env")
	}
	return config, nil
}

// getEnv retorna o valor de uma variável de ambiente ou um valor padrão
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetDNS retorna a string de conexão do banco de dados
func (c *DatabaseConfig) GetDSN() string {
    return fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
    )
}