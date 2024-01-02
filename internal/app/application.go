package app

import (
	"TraiveTest/internal/models"
	"TraiveTest/internal/repositories"
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// ListTransactionsResponse represents the response format for the ListTransactions endpoint.
type ListTransactionsResponse struct {
	Data         []models.Transaction `json:"data"`
	TotalRecords int                  `json:"totalRecords"`
	Page         int                  `json:"page"`
	PageSize     int                  `json:"pageSize"`
}

// Application represents the application logic.
type Application struct {
	TransactionRepository *repositories.TransactionRepository
	RabbitMQConfig        RabbitMQConfig
}

// NewApplication creates a new instance of Application.
func NewApplication(transactionRepository *repositories.TransactionRepository, rabbitMQ RabbitMQConfig) *Application {
	return &Application{
		TransactionRepository: transactionRepository,
		RabbitMQConfig:        rabbitMQ,
	}
}

// validates the input data for creating a transaction.
func ValidateTransactionInput(input models.TransactionValidation) error {
	validator := validator.New()

	return validator.Struct(input)
}

// CreateTransaction creates a new transaction.
func (app *Application) CreateTransaction(transaction *models.Transaction) error {

	//validate input data here before creating the transaction.
	if err := ValidateTransactionInput(models.NewTransactionValidation(transaction)); err != nil {
		return err
	}

	json, _ := json.Marshal(transaction)
	// add message notification on RabbitMQ
	app.RabbitMQConfig.PublishNotification(string(json))

	//business logic or validation.
	return app.TransactionRepository.CreateTransaction(transaction)
}

// ListTransactions retrieves a list of transactions with pagination and filtering.
func (app *Application) ListTransactions(filter map[string]interface{}, page, pageSize int) (ListTransactionsResponse, error) {
	transactions, err := app.TransactionRepository.ListTransactions(filter, page, pageSize)

	if err != nil {
		return ListTransactionsResponse{}, err
	}

	// Create the response
	response := ListTransactionsResponse{
		Data:         transactions,
		TotalRecords: len(transactions),
		Page:         page,
		PageSize:     pageSize,
	}

	return response, nil
}
