package internal

import (
	"TraiveTest/internal/app"
	"TraiveTest/internal/handlers"
	"TraiveTest/internal/repositories"
	"TraiveTest/internal/server"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SetupServer initializes and configures the server.
func SetupServer() (*server.Server, error) {
	// load env variables just once in here so can be use in any other place
	//config.InitEnvConfigs()

	// Initialize MongoDB connection
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://mongodb:27017"))
	//mongoClient, err := SetupMongoDB()
	if err != nil {
		return nil, err
	}

	db := mongoClient.Database("mytransactionsdb")

	// Initialize repositories
	transactionRepository := repositories.NewTransactionRepository(db)

	// Load RabbitMQ configuration from your application settings
	rabbitMQConfig := app.RabbitMQConfig{
		URL:      "amqp://admin:123456@rabbitmq:5672/",
		Exchange: "notificationExchange",
		Queue:    "notification",
	}

	// Initialize services
	transactionService := app.NewApplication(transactionRepository, rabbitMQConfig)

	// Initialize handlers
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Create and configure the server
	server := server.NewServer(transactionHandler)

	return server, nil
}

func SetupMongoDB() (*mongo.Client, error) {
	// MongoDB connection string with username and password
	connectionString := "mongodb://appUser:123456@mongodb:27017/"

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
