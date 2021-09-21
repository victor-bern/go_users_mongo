package server

import "github.com/gin-gonic/gin"

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

}
