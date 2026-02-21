package update

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	ID     int `json:"id" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

func (h *Handler) UpdateTransaction(c *gin.Context) {
	var requestBody RequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecaseUpdateTransaction.Update(c.Request.Context(), requestBody.ID, requestBody.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  http.StatusOK,
		"value":   fmt.Sprintf("transaction amount was changed to %.2f", requestBody.Amount)})
}
