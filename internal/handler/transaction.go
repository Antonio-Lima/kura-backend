package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kura/internal/auth"
	"kura/internal/database"
	"kura/internal/model"
	"kura/internal/service"
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
			"date":        newTransaction.Date,
		},
	})
}

func UpdateTransaction(c *gin.Context) {
	userID := auth.GetUserID(c)
	transactionID := c.Param("id")

	var input model.TransactionInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	transaction, isValid := service.FindAndValidateTransaction(c, userID, transactionID)

	if !isValid {
		return
	}

	transaction.Amount = input.Amount
	transaction.Description = input.Description
	transaction.CategoryId = input.CategoryId
	transaction.Date = input.Date
	transaction.Type = input.Type

	if err := database.DB.Save(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao atualizar lançamento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lançamento atualizado com sucesso",
		"data": gin.H{
			"amount":      transaction.Amount,
			"description": transaction.Description,
			"category_id": transaction.CategoryId,
			"date":        transaction.Date,
			"type":        transaction.Type,
		},
	})
}

func DeleteTransaction(c *gin.Context) {
	userID := auth.GetUserID(c)
	transactionID := c.Param("id")

	_, isValid := service.FindAndValidateTransaction(c, userID, transactionID)

	if !isValid {
		return
	}

	result := database.DB.Where("user_id = ? AND id = ?", userID, transactionID).Delete(&model.Transaction{})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao deletar lançamento",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Lançamento não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lançamento deletado com sucesso",
	})
}
