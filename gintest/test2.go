package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "List of users"})
		})
		v1.POST("/posts", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "List of posts"})
		})
	}
	r.Run(":8080")
}
