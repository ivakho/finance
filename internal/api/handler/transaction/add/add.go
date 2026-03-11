package add

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	CategoryID int `json:"category_id" binding:"required"`
	Amount     int64 `json:"amount" binding:"required"`
}

func (h *Handler) AddTransaction(c *gin.Context) {
	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecaseTransactionAdd.Add(c.Request.Context(), requestBody.CategoryID, requestBody.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "status": http.StatusOK})
}
