package get_expense_total

import (
	"finance/internal/api/handler/category"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestParams struct {
	DateFrom string `form:"date_from"`
	DateTo   string `form:"date_to"`
}

func (h *Handler) GetCategoryExpenseTotal(c *gin.Context) {
	var req RequestParams
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.DateFrom == "" || req.DateTo == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "date_from and date_to are required",
		})
		return
	}

	var dateFrom, dateTo time.Time
	dateFromParsed, err := time.Parse(time.DateOnly, req.DateFrom)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dateFrom = dateFromParsed

	dateToParsed, err := time.Parse(time.DateOnly, req.DateTo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dateTo = dateToParsed

	result, err := h.usecaseGetExpenseTotal.GetCategoryExpenseTotal(c.Request.Context(), dateFrom, dateTo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	categories := make([]category.CategoryTotal, 0, len(result))

	for _, v := range result {
		categories = append(categories, category.CategoryTotal{
			ID:    v.ID,
			Name:  v.Name,
			Total: v.Total,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  http.StatusOK,
		"value":   categories,
	})
}
