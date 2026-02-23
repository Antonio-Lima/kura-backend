package handler

import (
	"kura/internal/database"
	"kura/internal/model"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var input model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao criptografar a senha",
			"details": err.Error(),
		})
		return
	}
	input.Password = string(passwordHash)

	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "E-mail já cadastrado",
		})
		return
	}

	if err := database.SetupInitialUserData(input.ID); err != nil {
		println("Erro ao criar categorias padrão")
	}

	input.Password = ""

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso",
		"data":    input,
	})
}
