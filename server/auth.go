package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func authenticateReq(c *gin.Context) {
	k, exist := c.Request.Header["Secret"]
	if !exist {
		c.JSON(401, gin.H{"error": "No secretKey in cookie"})
		c.Abort()
		return
	}

	key := os.Getenv("SECRET")
	log.Println(key)
	if k[0] != key {
		c.JSON(401, gin.H{"error": "Invalid secretKey"})
		c.Abort()
		return
	}

	c.Next()
}
