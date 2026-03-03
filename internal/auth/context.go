package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserID(c *gin.Context) uuid.UUID {
	idStr, _ := c.Get("userID")
	userID, _ := uuid.Parse(fmt.Sprintf("%v", idStr))
	return userID
}
