package routes

import (
	"example/api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AlbumRoutes(r *gin.Engine) {
	r.GET("/albums", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, models.Albums)
	})

	r.POST("addAlbum", func(ctx *gin.Context) {

		var newAlbum models.Album

		err := ctx.BindJSON(&newAlbum)
		if err != nil {
			ctx.String(400, fmt.Sprint(err))
			return
		}
		newAlbum.ID = fmt.Sprint(len(models.Albums) + 1)
		models.Albums = append(models.Albums, newAlbum)
		ctx.IndentedJSON(200, gin.H{
			"message":  "album added succesfully",
			"new_abum": newAlbum,
		})
	})

	r.GET("album/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for _, album := range models.Albums {
			if album.ID == id {
				ctx.IndentedJSON(200, album)
				return
			}
		}
		ctx.String(400, "album not found")
	})

	r.DELETE("album/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		for _, album := range models.Albums {
			if album.ID == id {
				ctx.String(200, "album deleted succesfully")
				return
			}
		}
		ctx.String(400, "album not found")

	})
}
