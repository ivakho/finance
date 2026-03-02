package get

import (
	"finance/internal/storage/transaction"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestParams struct {
	CategoryID int     `form:"category_id"`
	Type       string  `form:"type"`
	DateFrom   *string `form:"date_from"`
	DateTo     *string `form:"date_to"`
	Limit      int     `form:"limit"`
}

func (h *Handler) GetTransaction(c *gin.Context) {
	var req RequestParams
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dateFrom, dateTo *time.Time
	if req.DateFrom != nil {
		time, _ := time.Parse("2006-01-02", *req.DateFrom)
		dateFrom = &time
	}

	if req.DateTo != nil {
		time, _ := time.Parse("2006-01-02", *req.DateTo)
		dateTo = &time
	}

	filter := transaction.TransactionFilter{
		CategoryID: req.CategoryID,
		Type:       req.Type,
		DateFrom:   dateFrom,
		DateTo:     dateTo,
		Limit:      req.Limit,
	}

	transactions, err := h.usecaseGetTransaction.Get(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"value": transactions})
}
