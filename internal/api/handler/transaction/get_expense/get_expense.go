package get_expense

import (
	"finance/internal/api/handler/transaction"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetExpense(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.usecaseGetExpense.GetExpense(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	transactions := make(transaction.Transactions, 0, len(result.Value))
	for _, v := range result.Value {
		transactions = append(transactions, transaction.Transaction{
			ID:        v.ID,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"value": transactions, "total": result.Total})
}
