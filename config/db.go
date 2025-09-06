package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}
}

func ConnectDatabase() {
	LoadEnv()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL não encontrada nas variáveis de ambiente")
	}

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	database, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal("Erro ao obter database SQL:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	DB = database
	log.Println("✅ Banco de dados conectado com sucesso!")
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database não inicializado. Chame ConnectDatabase() primeiro.")
	}
	return DB
}
