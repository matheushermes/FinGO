package main

import (
	"fmt"
	"time"

	_ "github.com/matheushermes/FinGO/init"
	"github.com/matheushermes/FinGO/internal/cache"
	"github.com/matheushermes/FinGO/internal/server"
	"github.com/matheushermes/FinGO/internal/worker"
)

func main() {
	fmt.Println("Iniciando projeto FinGO!")
	cache.InitRedis()
	worker.StartCryptoUpdater(5 * time.Minute)
	server := server.NewServer()
	server.RunServer()
}