package handler

import (
	"kura/internal/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var input model.CategoryInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	newCategory := model.Category{
		ID:            123,
		UserId:        1,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		CategoryInput: input,
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Categoria cadastrada com sucesso",
		"data":    newCategory,
	})
}
