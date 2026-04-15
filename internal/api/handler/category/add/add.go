package add

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestBody struct {
	Name string `json:"name" binding:"required"`
}

func (h *Handler) AddCategory(c *gin.Context) {
	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(requestBody.Name) > 30 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("too long category name")})
		return
	}

	if len(requestBody.Name) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("too short category name")})
		return
	}

	if err := h.usecaseCategoryAdd.Add(c.Request.Context(), requestBody.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "status": http.StatusOK})
}
