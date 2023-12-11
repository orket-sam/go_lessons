package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/orket-sam/go_lessons/controllers"
)

func AlbumRoutes(r *gin.Engine) {
	r.GET("/", controllers.GetAlbums)
	r.POST("/addAlbum", controllers.PostAlbum)
	r.GET("/album", controllers.GetAlbumById)
}
