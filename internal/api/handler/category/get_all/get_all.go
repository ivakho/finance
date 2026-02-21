package get_all

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllCategory(c *gin.Context) {

	categories, err := h.usecaseGetAllCategory.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "status": http.StatusOK, "value": categories})
}
