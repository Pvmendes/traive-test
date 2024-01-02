package repositories

import (
	"testing"
	"time"

	"TraiveTest/internal/models"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateTransaction(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("test", func(mt *mtest.T) {

		mt.AddMockResponses(mtest.CreateSuccessResponse(bson.D{{Key: "ID", Value: 2}}...))

		trans := NewTransactionRepository(mt.DB)
		testTransaction := models.NewTransaction(
			"2",
			"test-origin",
			"1001",
			50.0,
			"credit",
			time.Now(),
		)

		err := trans.CreateTransaction(testTransaction)

		assert.Nil(t, err)
	})
}

func TestCreateTransaction_fail(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("test", func(mt *mtest.T) {

		trans := NewTransactionRepository(mt.DB)
		testTransaction := new(models.Transaction)

		err := trans.CreateTransaction(testTransaction)

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

		killCursors := mtest.CreateCursorResponse(0,"DBName.CollectionName",mtest.NextBatch)

		mt.AddMockResponses(find, getMore, killCursors)

		trans := NewTransactionRepository(mt.DB)

		result, _ := trans.ListTransactions(filter, page, pageSize)

		assert.Equal(t, 2, len(result))
		assert.Equal(t, "1", result[0].ID)
		assert.Equal(t, "2", result[1].ID)
	})
}
