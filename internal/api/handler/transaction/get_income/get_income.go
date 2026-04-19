package get_income

import (
	"finance/internal/api/handler/transaction"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestParams struct {
	CategoryID int    `form:"category_id"`
	DateFrom   string `form:"date_from"`
	DateTo     string `form:"date_to"`
}

func (h *Handler) GetIncome(c *gin.Context) {
	var req RequestParams
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dateFrom, dateTo time.Time
	if req.DateFrom != "" {
		time, err := time.Parse(time.DateOnly, req.DateFrom)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dateFrom = time
	}

	if req.DateTo != "" {
		time, err := time.Parse(time.DateOnly, req.DateTo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dateTo = time
	}

	result, err := h.usecaseGetIncome.GetIncome(c.Request.Context(), req.CategoryID, dateFrom, dateTo)
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
