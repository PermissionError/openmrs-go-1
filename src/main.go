package main

import "github.com/gin-gonic/gin"
import "net/http"

const NameParam = "name"

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
	r.GET("/greet/:" + NameParam, func(c *gin.Context) {
		name := c.Param(NameParam)
		message := "Hello " + name + "!"
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
