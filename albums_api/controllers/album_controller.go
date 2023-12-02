package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/orket-sam/go_lessons/config"
	"github.com/orket-sam/go_lessons/models"
)

var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
	var albums = []models.Album{}
	config.DB.Find(&albums)
	if len(albums) == 0 {
		c.IndentedJSON(200, gin.H{
			"message": "No albums available",
		})
	} else {
		c.IndentedJSON(200, albums)
	}

}

func PostAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.String(200, newAlbum.Title+" Added!!")
}
