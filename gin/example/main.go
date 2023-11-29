package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": http.StatusAccepted,
		})
	})
	r.Run(":7890")
}
