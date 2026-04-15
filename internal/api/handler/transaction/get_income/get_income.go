package get_income

import (
	"finance/internal/api/handler/transaction"
	// "log"
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
		// log.Printf("ERROR: Failed to bind query: %v", err)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// log.Printf("Request params: category_id=%d, date_from=%s, date_to=%s",
	// 	req.CategoryID, req.DateFrom, req.DateTo)

	var dateFrom, dateTo time.Time
	if req.DateFrom != "" {
		time, err := time.Parse(time.DateOnly, req.DateFrom)
		// log.Printf("ERROR: Failed to parse date_from: %v", err)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dateFrom = time

		// log.Printf("Parsed date_from: %v", dateFrom)
	}

	if req.DateTo != "" {
		time, err := time.Parse(time.DateOnly, req.DateTo)
		// log.Printf("ERROR: Failed to parse date_to: %v", err)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dateTo = time

		// log.Printf("Parsed date_to: %v", dateTo)

	}

	result, err := h.usecaseGetIncome.GetIncome(c.Request.Context(), req.CategoryID, dateFrom, dateTo)
	if err != nil {

		// log.Printf("ERROR: Usecase failed: %v", err)

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
