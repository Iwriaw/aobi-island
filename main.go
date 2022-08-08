package main

import (
	"net/http"

	"github.com/Iwriaw/aobi-island/gin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Register(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:80")
}
