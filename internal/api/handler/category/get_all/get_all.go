package get_all

import (
	"finance/internal/api/handler/category"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllCategory(c *gin.Context) {

	result, err := h.usecaseGetAllCategory.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	categories := make([]category.Category, 0, len(result))

	for _, v := range result {
		categories = append(categories, category.Category{ID: v.ID, Name: v.Name, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt})
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "status": http.StatusOK, "value": categories})
}
