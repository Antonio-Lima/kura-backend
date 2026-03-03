package handler

import (
	"kura/internal/auth"
	"kura/internal/database"
	"kura/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	userID := auth.GetUserID(c)

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

func GetCategoryByID(c *gin.Context) {
	userID := auth.GetUserID(c)

	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da categoria inválido",
		})
		return
	}

	var category model.Category

	if err := database.DB.Where("user_id = ? AND id = ?", userID, categoryID).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Categoria não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Categoria encontrada com sucesso",
		"data": gin.H{
			"id":       category.ID,
			"category": category.Category,
			"color":    category.Color,
			"icon":     category.Icon,
		},
	})
}

func CreateCategory(c *gin.Context) {
	userID := auth.GetUserID(c)

	var input model.CategoryInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	newCategory := model.Category{
		UserID:        userID,
		CategoryInput: input,
	}

	if err := database.DB.Create(&newCategory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao cadastrar categoria",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Categoria cadastrada com sucesso",
		"data": gin.H{
			"id":       newCategory.ID,
			"category": newCategory.Category,
			"color":    newCategory.Color,
			"icon":     newCategory.Icon,
		},
	})
}

func UpdateCategory(c *gin.Context) {
	userID := auth.GetUserID(c)

	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da categoria inválido",
		})
		return
	}

	var input model.CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	var category model.Category
	if err := database.DB.Where("user_id = ? AND id = ?", userID, categoryID).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Categoria não encontrada",
		})
		return
	}

	category.Category = input.Category
	category.Color = input.Color
	category.Icon = input.Icon

	if err := database.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao atualizar categoria",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Categoria atualizada com sucesso",
		"data": gin.H{
			"id":         category.ID,
			"category":   category.Category,
			"color":      category.Color,
			"icon":       category.Icon,
			"updated_at": category.UpdatedAt,
		},
	})
}

func DeleteCategory(c *gin.Context) {
	userId := auth.GetUserID(c)

	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da categoria inválido",
		})
		return
	}

	var category model.Category
	if err := database.DB.Where("user_id = ? AND id = ?", userId, categoryID).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Categoria não encontrada",
		})
		return
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao deletar categoria",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Categoria deletada com sucesso",
	})
}
