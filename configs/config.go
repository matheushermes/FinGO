package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT              = ""
	POSTGRES_HOST      = ""
	POSTGRES_USER      = ""
	POSTGRES_PASSWORD  = ""
	POSTGRES_DB        = ""
	POSTGRES_PORT      = ""
	STRING_CONNECTION  = ""
)

func LoadingEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Erro ao carregar arquivo .env: ", err)
	}

	PORT = os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("A variável de ambiente PORT não está definida no arquivo .env")
	}

	POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	if POSTGRES_HOST == "" {
		log.Fatal("A variável de ambiente POSTGRES_HOST não está definida no arquivo .env")
	}

	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	if POSTGRES_USER == "" {
		log.Fatal("A variável de ambiente POSTGRES_USER não está definida no arquivo .env")
	}
	
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	if POSTGRES_PASSWORD == "" {
		log.Fatal("A variável de ambiente POSTGRES_PASSWORD não está definida no arquivo .env")
	}

	POSTGRES_DB = os.Getenv("POSTGRES_DB")
	if POSTGRES_DB == "" {
		log.Fatal("A variável de ambiente POSTGRES_DB não está definida no arquivo .env")
	}

	POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	if POSTGRES_PORT == "" {
		log.Fatal("A variável de ambiente POSTGRES_PORT não está definida no arquivo .env")
	}

	STRING_CONNECTION = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

}