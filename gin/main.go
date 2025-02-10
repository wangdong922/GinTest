package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"log"
	"net/http"
)

// 自定义中间件
func MyHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("username", "zhangsan")
		context.Next()
		//context.Abort()
	}
}
func main() {
	ginServer := gin.Default()

	ginServer.Use(favicon.New("./1.jpg"))
	ginServer.Use(MyHandler())
	ginServer.LoadHTMLGlob("template/*")
	ginServer.Static("/static", "./static")

	ginServer.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})
	ginServer.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"message": "hello world"})
	})

	ginServer.GET("user/info", MyHandler(), func(c *gin.Context) {
		stringsession := c.MustGet("username")
		log.Println("=================>", stringsession)
		userId := c.Query("userid")
		username := c.Query("username")
		c.JSON(http.StatusOK, gin.H{
			"userId":   userId,
			"username": username,
		})
	})

	ginServer.GET("user/list/:userid/:username", func(c *gin.Context) {
		userid := c.Param("userid")
		username := c.Param("username")
		c.JSON(http.StatusOK, gin.H{
			"userId":   userid,
			"username": username,
		})
	})

	ginServer.POST("/json", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(data, &m)
		c.JSON(http.StatusOK, m)
	})

	ginServer.POST("/user/add", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	// 路由  重定向
	ginServer.GET("test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.kuangstudy.com")
	})

	ginServer.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	ginServer.Run(":8080")
}
