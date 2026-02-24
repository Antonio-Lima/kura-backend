package handler

import (
	"kura/internal/auth"
	"kura/internal/database"
	"kura/internal/model"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
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
	user := model.User{
		UserBase: input.UserBase,
		Password: string(passwordHash),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "E-mail já cadastrado",
		})
		return
	}

	if err := database.SetupInitialUserData(user.ID); err != nil {
		println("Erro ao criar categorias padrão")
	}

	user.Password = ""

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso",
		"data":    user,
	})
}

func LoginUser(c *gin.Context) {
	var input model.UserLogin

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user model.User
	if err := database.DB.Where("email= ?", input.Email).First(&user).Error; err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "E-mail ou senha inválidos",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "E-mail ou senha inválidos",
		})
		return
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao gerar token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login realizado com sucesso",
		"token":   token,
	})
}
