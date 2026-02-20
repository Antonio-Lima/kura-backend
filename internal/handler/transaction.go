package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"kura/internal/model"
)

func CreateTransaction(c *gin.Context) {
	var input model.TransactionInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	newTransaction := model.Transaction{
		ID:               123,
		UserId:           1,
		CreatedAt:        time.Now(),
		TransactionInput: input,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transação validada com sucesso!",
		"data":    newTransaction,
	})
}
