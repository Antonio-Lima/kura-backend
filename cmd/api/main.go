package main

import (
	"kura/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/transactions", handler.CreateTransaction)

	r.POST("/category", handler.CreateCategory)

	r.POST("/user", handler.CreateUser)
	r.GET("/user", handler.GetUser)

	r.Run(":8080")
}
