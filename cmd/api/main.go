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
		// CATEGORY ROUTES
		protected.POST("/categories", handler.CreateCategory)
		protected.GET("/categories", handler.GetCategories)
		protected.GET("/categories/:id", handler.GetCategoryByID)
		protected.PUT("/categories/:id", handler.UpdateCategory)
		protected.DELETE("/categories/:id", handler.DeleteCategory)
		// CATEGORY ROUTES

		// TRANSACTION ROUTES
		protected.POST("/transactions", handler.CreateTransaction)
		protected.PUT("/transactions/:id", handler.UpdateTransaction)
		protected.DELETE("/transactions/:id", handler.DeleteTransaction)
		// TRANSACTION ROUTES
	}

	r.Run(":8080")
}
