package database

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

// ConectaComBancoDeDados initializes the database connection
func ConectaComBancoDeDados() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao carregar arquivo .env, usando variáveis de ambiente do sistema.")
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")
	sslmode := os.Getenv("SSL_MODE")

	if host == "" || user == "" || password == "" || dbname == "" || port == "" || sslmode == "" {
		log.Panic("Um ou mais parâmetros de conexão com o banco de dados estão vazios.")
	}

	stringDeConexao := "host=" + host +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" port=" + port +
		" sslmode=" + sslmode

	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados: ", err)
	}

	if err := DB.AutoMigrate(&models.Aluno{}); err != nil {
		log.Panic("Erro ao realizar a migração: ", err)
	}
}

// FechaConexaoComBancoDeDados closes the database connection
func FechaConexaoComBancoDeDados() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Panic("Erro ao obter a instância de sql.DB: ", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Panic("Erro ao fechar a conexão com banco de dados: ", err)
	}
}
