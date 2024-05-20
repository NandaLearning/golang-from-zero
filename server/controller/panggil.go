package controller

import "github.com/gin-gonic/gin"

func Panggil(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "nanda",
	})
}
