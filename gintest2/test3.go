package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/aaa", func(c *gin.Context) {
		var json struct {
			Name  string
			Email string
		}
		if err := c.ShouldBind(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name":  json.Name,
			"email": json.Email,
		})
	})
	r.Run(":8080")
}
