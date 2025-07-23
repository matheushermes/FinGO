package main

import (
	"fmt"

	_ "github.com/matheushermes/FinGO/init"
	"github.com/matheushermes/FinGO/internal/server"
)

func main() {
	fmt.Println("Iniciando projeto FinGO!")
	server := server.NewServer()
	server.RunServer()
}