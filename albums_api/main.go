package main

import (
	"example/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	routes.AlbumRoutes(r)
	r.Run("localhost:8080")
}
