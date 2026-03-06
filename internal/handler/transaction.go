package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kura/internal/auth"
	"kura/internal/database"
	"kura/internal/model"
)

func CreateTransaction(c *gin.Context) {
	userID := auth.GetUserID(c)
	var input model.TransactionInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	newTransaction := model.Transaction{
		UserId:           userID,
		TransactionInput: input,
	}

	if err := database.DB.Where("user_id = ? AND id = ?", userID, newTransaction.CategoryId).First(&model.Category{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Categoria inválida",
		})
		return
	}

	if err := database.DB.Create(&newTransaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Ocorreu um erro ao cadastrar transação",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transação criada com sucesso!",
		"data": gin.H{
			"id":          newTransaction.ID,
			"description": newTransaction.Description,
			"amount":      newTransaction.Amount,
			"type":        newTransaction.Type,
			"category":    newTransaction.CategoryId,
		},
	})
}
