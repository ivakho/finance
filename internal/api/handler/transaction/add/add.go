package add

import (
	"net/http"
	"time"

	"finance/internal/usecase/transaction"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	CategoryID int    `json:"category_id" binding:"required"`
	Type       string `json:"type" binding:"required"`
	Amount     int64  `json:"amount" binding:"required"`
	CreatedAt  string `json:"created_at" binding:"required"`
}

func (h *Handler) AddTransaction(c *gin.Context) {
	var req RequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var createdAt time.Time
	if req.CreatedAt != "" {
		time, err := time.Parse(time.DateOnly, req.CreatedAt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		createdAt = time
	}

	tx := transaction.TransactionAdd{
		CategoryID: req.CategoryID,
		TxType:     req.Type,
		Amount:     req.Amount,
		CreatedAt:  createdAt,
	}

	if err := h.usecaseTransactionAdd.Add(c.Request.Context(), tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
