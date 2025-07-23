package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/matheushermes/FinGO/configs"
	"github.com/matheushermes/FinGO/internal/server/routes"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer() Server {
	return Server {
		port: configs.PORT,
		server: gin.Default(),
	}
}

func (s *Server) RunServer() {
	router := routes.ConfigRoutes(s.server)
	fmt.Println("Starting server on port", s.port)
	log.Fatal(router.Run(":" + s.port))
}