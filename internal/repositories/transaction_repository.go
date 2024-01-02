package repositories

import (
	"TraiveTest/internal/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TransactionRepository represents the repository for managing transactions in MongoDB.
type TransactionRepository struct {
	collection *mongo.Collection
}

// NewTransactionRepository creates a new instance of TransactionRepository.
func NewTransactionRepository(db *mongo.Database) *TransactionRepository {
	collection := db.Collection("transactions")
	return &TransactionRepository{
		collection: collection,
	}
}

// CreateTransaction inserts a new transaction record into the MongoDB collection.
func (repo *TransactionRepository) CreateTransaction(transaction *models.Transaction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := repo.collection.InsertOne(ctx, transaction)
	if err != nil {
		log.Printf("Failed to create transaction: %v\n", err)
		return err
	}

	return nil
}

// ListTransactions retrieves a list of transactions from the MongoDB collection.
func (repo *TransactionRepository) ListTransactions(filter bson.M, page, pageSize int) ([]models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * pageSize))
	findOptions.SetLimit(int64(pageSize))

	cursor, err := repo.collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Printf("Failed to list transactions: %v\n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var transactions []models.Transaction
	if err := cursor.All(ctx, &transactions); err != nil {
		log.Printf("Failed to decode transactions: %v\n", err)
		return nil, err
	}

	return transactions, nil
}
