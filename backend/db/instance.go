package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Repo *gorm.DB

func ConnectDatabaseGorm() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if user == "" || password == "" || host == "" || port == "" || dbName == "" {
		log.Fatal("Variáveis de ambiente do banco de dados não configuradas corretamente")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)

	engine, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados: ", err)
	}

	sqlDB, err := engine.DB()
	if err != nil {
		log.Fatal("Falha ao obter a conexão do banco de dados: ", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Falha ao pingar o banco de dados: ", err)
	}

	Repo = engine

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso")

	return Repo // Retorna a instância do banco de dados
}

func Migrate(engine *gorm.DB) error {
	tables := []interface{}{}

	for _, table := range tables {
		if err := engine.AutoMigrate(table); err != nil {
			return fmt.Errorf("falha ao migrar a tabela %T: %v", table, err)
		}
	}
	return nil
}
