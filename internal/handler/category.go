package handler

import (
	"fmt"
	"kura/internal/database"
	"kura/internal/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		ID:            1,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		UserID:        uuid.New(),
		CategoryInput: input,
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Categoria cadastrada com sucesso",
		"data":    newCategory,
	})
}

func GetCategories(c *gin.Context) {
	idStr, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao recuperar usuário",
		})
	}

	userID, _ := uuid.Parse(fmt.Sprintf("%v", idStr))

	var categories []model.Category

	if err := database.DB.Where("user_id = ?", userID).Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao buscar categorias",
			"error":   err.Error(),
		})
		return
	}

	type Response struct {
		ID       uint   `json:"id"`
		Category string `json:"category"`
		Color    string `json:"color"`
		Icon     string `json:"icon"`
	}

	var result []Response
	for _, cat := range categories {
		result = append(result, Response{
			ID:       cat.ID,
			Category: cat.Category,
			Color:    cat.Color,
			Icon:     cat.Icon,
		})
	}

	c.JSON(http.StatusOK, result)
}
