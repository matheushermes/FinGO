package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT = ""
)

func LoadingEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Erro ao carregar arquivo .env: ", err)
	}

	PORT = os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("A variável de ambiente PORT não está definida no arquivo .env")
	}
}