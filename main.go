package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.StaticFile("/", "index.html")

	count := 0

	r.GET("/count", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"count": count,
		})
	})

	r.POST("/count", func(c *gin.Context) {
		count++
		c.JSON(http.StatusOK, gin.H{
			"count": count,
		})
	})

	r.Run(":8080")
}
