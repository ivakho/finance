package get_income

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	ID        int       `json:"id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Transactions []Transaction

func (h *Handler) GetIncome(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.usecaseGetIncome.GetIncome(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	transactions := make(Transactions, 0, len(result.Value))
	for _, v := range result.Value {
		transactions = append(transactions, Transaction{ID: v.ID, Amount: v.Amount, CreatedAt: v.CreatedAt})
	}

	c.JSON(http.StatusOK, gin.H{"value": transactions, "total": result.Total})
}
