package server

import (
	"gomongo/src/server/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   ":5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	routes := routes.ConfigRoutes(s.server)

	log.Print("Server is running on port ", s.port)
	routes.Run(s.port)

}
