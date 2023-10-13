package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}