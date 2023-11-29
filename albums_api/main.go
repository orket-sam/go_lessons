package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type albums struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var allALbums = []albums{
	{ID: "12", Title: "Life of Pablo", Artist: "Kanye West", Price: 21.23},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	r := gin.Default()
	r.GET("/api/getAlbums", getAlbums)
	r.POST("api/getAlbums", postAlbums)
	r.Run("localhost:8000")

}

func getAlbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, allALbums)
}

func postAlbums(c *gin.Context) {
	var newAlbum albums

	err := c.BindJSON(&newAlbum)
	if err != nil {
		print("oops")
		return
	}

	// Add the new album to the slice.
	allALbums = append(allALbums, newAlbum)
	c.IndentedJSON(http.StatusCreated, allALbums)
}
