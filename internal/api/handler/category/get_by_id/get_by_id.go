package get_by_id

import (
	"finance/internal/api/handler/category"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.usecaseGetCategoryByID.GetCategoryByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	categoryByID := category.Category{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  http.StatusOK,
		"value":   categoryByID,
	})
}
