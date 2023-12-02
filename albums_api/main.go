package main

import (
	"github.com/gin-gonic/gin"
	"github.com/orket-sam/go_lessons/config"
	"github.com/orket-sam/go_lessons/routes"
)

func main() {
	r := gin.Default()
	config.ConnectDb()
	routes.AlbumRoutes(r)
	r.Run(":8080")
}
