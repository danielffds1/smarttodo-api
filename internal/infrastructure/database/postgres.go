package database

import (
	"fmt"
	"log"
	"time"

	"smarttodo-api/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewPostgresConnection cria uma instancia nova conexão com o PostgreSQL
func NewPostgresConnection(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	// Obtém a string de conexão do banco de dados
	dsn := cfg.GetDSN()

	// Configura o logger do GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	// Tentativa de conexão
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	// Config do pool de conexões
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("falha ao obter SQL DB: %w", err)
	}

	// SetMaxIdleConns define o número máximo de conexões ociosas
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns define o número máximo de conexões abertas
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime define o tempo máximo de vida de uma conexão
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Testa a conexão
    if err := sqlDB.Ping(); err != nil {
        return nil, fmt.Errorf("falha ao pingar o banco de dados: %w", err)
    }

	log.Println("Conexão com o banco de dados estabelecida com sucesso")

	return db, nil
}

// Close fecha a conexão com o banco de dados
func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}