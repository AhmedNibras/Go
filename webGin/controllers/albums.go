package controllers

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)


// albums slice to seed record data.
var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}


// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}


// GetAlbumsByID responds with the the album of the given id.
func GetAlbumsByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}


// PostAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {

	var newAlbum models.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}


// PutAlbums updates an album identified by the id.
func PutAlbums(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			var newAlbum models.Album
			if err := c.BindJSON(&newAlbum); err != nil {
				return
			}
			albums[i] = newAlbum
			c.IndentedJSON(http.StatusOK, newAlbum)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}


// DeleteAlbums deletes an album identified by the id.
func DeleteAlbums(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "album not found"})
}

// PatchAlbums partially updates an album identified by the id.
func PatchAlbums(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			var val models.Album
			if err := c.BindJSON(&val); err != nil {
				return
			}
			if val.Title != "" {
				albums[i].Title = val.Title
			}
			if val.Artist != "" {
				albums[i].Artist = val.Artist
			}
			if val.Price != 0 {
				albums[i].Price = val.Price
			}
			c.JSON(http.StatusOK, albums[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
