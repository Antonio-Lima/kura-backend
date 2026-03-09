package service

import (
	"kura/internal/database"
	"kura/internal/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FindAndValidateTransaction(c *gin.Context, userID uuid.UUID, transactionID string) (model.Transaction, bool) {
	var transaction model.Transaction

	if err := database.DB.Where("user_id = ? AND id = ?", userID, transactionID).First(&transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Lançamento não encontrado",
		})
		return transaction, false
	}

	editLimit := time.Now().AddDate(0, -1, 0)

	if transaction.Date.Before(editLimit) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Lançamento bloqueado para edição por ser posterior a 1 mês",
			"help":  "Para correções anteriores a 1 mês, realize um novo lançamento com o valor da diferença",
		})
		return transaction, false
	}

	return transaction, true
}
