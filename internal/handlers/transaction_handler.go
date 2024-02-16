package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"TraiveTest/internal/app"
	"TraiveTest/internal/models"
)

// TransactionHandler handles HTTP requests related to transactions.
type TransactionHandler struct {
	app *app.Application
}

// NewTransactionHandler creates a new instance of TransactionHandler.
func NewTransactionHandler(app *app.Application) *TransactionHandler {
	return &TransactionHandler{
		app: app,
	}
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Ping
// @Router /Ping [get]
func (handler *TransactionHandler) Ping(g *gin.Context) {
	g.JSON(http.StatusOK, "pong")
}

// @Summary Create a new transaction
// @Description CreateTransaction handles the creation of a new transaction
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param input body Transaction true "Transaction object"
// @Success 201 {object} Transaction
// @Router /transactions/ [post]
func (handler *TransactionHandler) CreateTransaction(c *gin.Context) {
	var input models.Transaction

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction := models.NewTransaction(
		input.ID,
		input.Origin,
		input.UserID,
		input.Amount,
		input.Operation,
		input.CreatedAt,
	)

	err := handler.app.CreateTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

// @Summary List transactions
// @Description List transactions with pagination and filtering
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Param origin query string false "Origin filter"
// @Param userID query string false "User ID filter"
// @Success 200 {array} Transaction
// @Router /transactions/ [get]
func (handler *TransactionHandler) ListTransactions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	origin := c.DefaultQuery("origin", "")
	userID := c.DefaultQuery("userID", "")

	filter := make(map[string]interface{})
	if origin != "" {
		filter["origin"] = origin
	}
	if userID != "" {
		filter["user_id"] = userID
	}

	transactions, err := handler.app.ListTransactions(filter, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
