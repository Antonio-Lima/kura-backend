package main

import (
	"kura/internal/database"
	"kura/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	r := gin.Default()

	// r.POST("/transactions", handler.CreateTransaction)

	r.POST("/user/register", handler.RegisterUser)

	r.Run(":8080")
}
