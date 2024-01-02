package models

import (
	"time"
)

// Transaction represents a monetary transaction made by a user.
// swagger:model
type Transaction struct {
	// The unique identifier for the transaction.
	//
	// required: true
	ID string `json:"_id,omitempty" bson:"_id,omitempty"`

	// The origin where the user started this transaction (e.g., "desktop-web", "mobile-android").
	//
	// required: true
	Origin string `json:"origin"`

	// The ID of the user that made the transaction.
	//
	// required: true
	UserID string `json:"user_id"`

	// The monetary amount of the transaction.
	//
	// required: true
	Amount float64 `json:"amount"`

	// The type of the operation, which can be either "credit" or "debit".
	//
	// required: true
	Operation string `json:"operation"`

	// The date and time when the transaction was created.
	//
	// required: true
	CreatedAt time.Time `json:"created_at"`
}

type TransactionValidation struct {
	ID        string
	Origin    string    `validate:"required,min=4,max=15"`
	UserID    string    `validate:"required"`
	Amount    float64   `validate:"required"`
	Operation string    `validate:"required,min=5,max=6"`
	CreatedAt time.Time `validate:"required"`
}

func NewTransactionValidation(transaction *Transaction) TransactionValidation {
	return TransactionValidation{
		ID:        transaction.ID,
		Origin:    transaction.Origin,
		UserID:    transaction.UserID,
		Amount:    transaction.Amount,
		Operation: transaction.Operation,
		CreatedAt: transaction.CreatedAt,
	}
}

// NewTransaction creates a new Transaction instance.
func NewTransaction(id, origin, userID string, amount float64, operation string, createdAt time.Time) *Transaction {
	return &Transaction{
		ID:        id,
		Origin:    origin,
		UserID:    userID,
		Amount:    amount,
		Operation: operation,
		CreatedAt: createdAt,
	}
}
