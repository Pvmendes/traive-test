package app

import (
	"testing"
	"time"

	"TraiveTest/internal/models"
	"TraiveTest/internal/repositories"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateTransaction(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("test", func(mt *mtest.T) {

		mt.AddMockResponses(mtest.CreateSuccessResponse(bson.D{{Key: "ID", Value: 2}}...))

		trans := repositories.NewTransactionRepository(mt.DB)

		rabbitMQConfig := RabbitMQConfig{
			URL:      "test",
			Exchange: "",
			Queue:    "",
		}

		transactionService := NewApplication(trans, rabbitMQConfig)

		testTransaction := models.NewTransaction(
			"2",
			"test-origin",
			"1001",
			50.0,
			"credit",
			time.Now(),
		)

		err := transactionService.CreateTransaction(testTransaction)

		assert.Nil(t, err)
	})
}

func TestCreateTransaction_Valid_ID_Required_Fail(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("test", func(mt *mtest.T) {

		trans := repositories.NewTransactionRepository(mt.DB)

		rabbitMQConfig := RabbitMQConfig{
			URL:      "test",
			Exchange: "",
			Queue:    "",
		}

		transactionService := NewApplication(trans, rabbitMQConfig)

		testTransaction := models.NewTransaction(
			"",
			"test-origin",
			"1001",
			50.0,
			"credit",
			time.Now(),
		)

		err := transactionService.CreateTransaction(testTransaction)

		assert.NotNil(t, err)
	})
}

func TestCreateTransaction_Valid_Origin_Required_Fail(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("test", func(mt *mtest.T) {

		trans := repositories.NewTransactionRepository(mt.DB)

		rabbitMQConfig := RabbitMQConfig{
			URL:      "test",
			Exchange: "",
			Queue:    "",
		}

		transactionService := NewApplication(trans, rabbitMQConfig)

		testTransaction := models.NewTransaction(
			"1",
			"",
			"1001",
			50.0,
			"credit",
			time.Now(),
		)

		err := transactionService.CreateTransaction(testTransaction)

		assert.NotNil(t, err)
	})
}

func TestCreateTransaction_Valid_UserID_Required_Fail(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("test", func(mt *mtest.T) {

		trans := repositories.NewTransactionRepository(mt.DB)

		rabbitMQConfig := RabbitMQConfig{
			URL:      "test",
			Exchange: "",
			Queue:    "",
		}

		transactionService := NewApplication(trans, rabbitMQConfig)

		testTransaction := models.NewTransaction(
			"1",
			"teste-origin",
			"",
			50.0,
			"credit",
			time.Now(),
		)

		err := transactionService.CreateTransaction(testTransaction)

		assert.NotNil(t, err)
	})
}

func TestCreateTransaction_Valid_Amount_Required_Fail(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("test", func(mt *mtest.T) {

		trans := repositories.NewTransactionRepository(mt.DB)

		rabbitMQConfig := RabbitMQConfig{
			URL:      "test",
			Exchange: "",
			Queue:    "",
		}

		transactionService := NewApplication(trans, rabbitMQConfig)

		testTransaction := models.NewTransaction(
			"2",
			"test-origin",
			"1001",
			0,
			"credit",
			time.Now(),
		)

		err := transactionService.CreateTransaction(testTransaction)

		assert.NotNil(t, err)
	})
}

func TestCreateTransaction_Valid_Operation_Required_Fail(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("test", func(mt *mtest.T) {

		trans := repositories.NewTransactionRepository(mt.DB)

		rabbitMQConfig := RabbitMQConfig{
			URL:      "test",
			Exchange: "",
			Queue:    "",
		}

		transactionService := NewApplication(trans, rabbitMQConfig)

		testTransaction := models.NewTransaction(
			"2",
			"test-origin",
			"1001",
			50.0,
			"",
			time.Now(),
		)

		err := transactionService.CreateTransaction(testTransaction)

		assert.NotNil(t, err)
	})
}

func TestCreateTransaction_Valid_Origin_Min_Fail(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("test", func(mt *mtest.T) {

		trans := repositories.NewTransactionRepository(mt.DB)

		rabbitMQConfig := RabbitMQConfig{
			URL:      "test",
			Exchange: "",
			Queue:    "",
		}

		transactionService := NewApplication(trans, rabbitMQConfig)

		testTransaction := models.NewTransaction(
			"2",
			"tes",
			"1001",
			50.0,
			"Credit",
			time.Now(),
		)

		err := transactionService.CreateTransaction(testTransaction)

		assert.NotNil(t, err)
	})
}

func TestCreateTransaction_Valid_Origin_Max_Fail(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("test", func(mt *mtest.T) {

		trans := repositories.NewTransactionRepository(mt.DB)

		rabbitMQConfig := RabbitMQConfig{
			URL:      "test",
			Exchange: "",
			Queue:    "",
		}

		transactionService := NewApplication(trans, rabbitMQConfig)

		testTransaction := models.NewTransaction(
			"2",
			"test-origin-test-kaz",
			"1001",
			50.0,
			"Credit",
			time.Now(),
		)

		err := transactionService.CreateTransaction(testTransaction)

		assert.NotNil(t, err)
	})
}

func TestCreateTransaction_Valid_Operation_Min_Fail(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("test", func(mt *mtest.T) {

		trans := repositories.NewTransactionRepository(mt.DB)

		rabbitMQConfig := RabbitMQConfig{
			URL:      "test",
			Exchange: "",
			Queue:    "",
		}

		transactionService := NewApplication(trans, rabbitMQConfig)

		testTransaction := models.NewTransaction(
			"2",
			"tes",
			"1001",
			50.0,
			"Cred",
			time.Now(),
		)

		err := transactionService.CreateTransaction(testTransaction)

		assert.NotNil(t, err)
	})
}

func TestCreateTransaction_Valid_Operation_Max_Fail(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("test", func(mt *mtest.T) {

		trans := repositories.NewTransactionRepository(mt.DB)

		rabbitMQConfig := RabbitMQConfig{
			URL:      "test",
			Exchange: "",
			Queue:    "",
		}

		transactionService := NewApplication(trans, rabbitMQConfig)

		testTransaction := models.NewTransaction(
			"2",
			"test-origin",
			"1001",
			50.0,
			"Credit-test-kaz",
			time.Now(),
		)

		err := transactionService.CreateTransaction(testTransaction)

		assert.NotNil(t, err)
	})
}

func TestListTransactions(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("test", func(mt *mtest.T) {

		page := 1
		pageSize := 10
		origin := ""
		userID := ""

		//validate and sanitize query parameters here.

		filter := make(map[string]interface{})
		if origin != "" {
			filter["origin"] = origin
		}
		if userID != "" {
			filter["user_id"] = userID
		}

		find := mtest.CreateCursorResponse(
			1,
			"DBName.CollectionName",
			mtest.FirstBatch,
			bson.D{
				{Key: "id", Value: "1"},
				{Key: "origin", Value: "origin"},
				{Key: "user_id", Value: "123"},
				{Key: "amount", Value: 14},
				{Key: "operation", Value: "credit"},
				{Key: "created_at", Value: time.Now()},
			})

		getMore := mtest.CreateCursorResponse(
			1,
			"DBName.CollectionName",
			mtest.NextBatch,
			bson.D{
				{Key: "id", Value: "2"},
				{Key: "origin", Value: "origin"},
				{Key: "user_id", Value: "321"},
				{Key: "amount", Value: 234},
				{Key: "operation", Value: "credit"},
				{Key: "created_at", Value: time.Now()},
			})

		killCursors := mtest.CreateCursorResponse(0, "DBName.CollectionName", mtest.NextBatch)

		mt.AddMockResponses(find, getMore, killCursors)

		trans := repositories.NewTransactionRepository(mt.DB)

		rabbitMQConfig := RabbitMQConfig{
			URL:      "test",
			Exchange: "",
			Queue:    "",
		}

		transactionService := NewApplication(trans, rabbitMQConfig)

		list, err := transactionService.ListTransactions(filter, page, pageSize)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(list.Data))
		assert.Equal(t, "1", list.Data[0].ID)
		assert.Equal(t, "2", list.Data[1].ID)
	})
}

func TestListTransactions_Fail(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("test", func(mt *mtest.T) {

		page := 1
		pageSize := 10
		origin := "origin"
		userID := ""

		filter := make(map[string]interface{})
		if origin != "" {
			filter["origin"] = origin
		}
		if userID != "" {
			filter["user_id"] = userID
		}

		trans := repositories.NewTransactionRepository(mt.DB)

		rabbitMQConfig := RabbitMQConfig{
			URL:      "test",
			Exchange: "",
			Queue:    "",
		}

		transactionService := NewApplication(trans, rabbitMQConfig)

		list, err := transactionService.ListTransactions(filter, page, pageSize)

		assert.NotNil(t, err)
		assert.Equal(t, 0, len(list.Data))
	})
}
