package server

import (
	"gomongo/src/server/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   ":" + os.Getenv("PORT"),
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	routes := routes.ConfigRoutes(s.server)

	log.Print("Server is running on port ", s.port)
	routes.Run(s.port)

}
