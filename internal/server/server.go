package server

import (
	"fmt"
	"log"

	"TraiveTest/internal/handlers"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server represents the HTTP server.
type Server struct {
	Router             *gin.Engine
	TransactionHandler *handlers.TransactionHandler
}

// NewServer creates a new instance of Server.
func NewServer(transactionHandler *handlers.TransactionHandler) *Server {
	server := &Server{
		Router:             gin.Default(),
		TransactionHandler: transactionHandler,
	}
	server.setupRoutes()
	return server
}

// setupRoutes configures the server's routes.
func (s *Server) setupRoutes() {

	transactionGroup := s.Router.Group("/transactions")
	{
		transactionGroup.POST("/CreateTransaction", s.TransactionHandler.CreateTransaction)
		transactionGroup.GET("/ListTransactions", s.TransactionHandler.ListTransactions)
	}

	pingBack := s.Router.Group("/ping")
	{
		pingBack.GET("/", s.TransactionHandler.Ping)
	}

	// Serve Swagger documentation
	s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// Run starts the HTTP server on the specified port.
func (s *Server) Run(port int) {

	addr := fmt.Sprintf(":%d", port)

	log.Printf("Server is running on http://localhost%s\n", addr)

	if err := s.Router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
