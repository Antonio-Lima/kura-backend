package handler

import (
	"kura/internal/model"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var input model.UserInput

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

	newUser := model.User{
		Base: model.Base{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UserBase: model.UserBase{
			Email:  input.Email,
			Name:   input.Name,
			Avatar: input.Avatar,
		},
		Password: string(passwordHash),
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso",
		"data":    newUser,
	})
}

func GetUser(c *gin.Context) {
	user := model.User{
		Base: model.Base{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UserBase: model.UserBase{
			Email:  "teste@teste.com",
			Name:   "Teste",
			Avatar: "https://example.com/avatar.png",
		},
		Password: "12345678",
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário encontrado",
		"data":    user,
	})
}
