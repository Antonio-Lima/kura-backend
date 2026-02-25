package main

import (
	"log"
	"os"

	"kura/internal/database"
	"kura/internal/handler"
	"kura/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("ERRO CRÍTICO: A variável JWT_SECRET não foi definida!")
	}

	database.Init()

	r := gin.Default()

	r.POST("/user/register", handler.RegisterUser)
	r.POST("/user/login", handler.LoginUser)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/categories", handler.GetCategories)
	}

	r.Run(":8080")
}
