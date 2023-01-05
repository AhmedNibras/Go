package controllers

import (
	"net/http"
	"test/models"
	"test/queries"

	"github.com/gin-gonic/gin"
)

// GetAlbums returns a list of all albums in the database.
func GetAlbums(c *gin.Context) {
	// Call the GetAlbumsQuery function to get a list of albums
	albums, err := queries.GetAlbumsQuery()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting albums"})
		return
	}
	// Return the list of albums as JSON
	c.JSON(http.StatusOK, albums)
}

// PostAlbums adds an album to the database.
func PostAlbums(c *gin.Context) {
	// Validate input
	var req models.Album
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Call the PostAlbumsQuery function to add the album
	album, err := queries.PostAlbumsQuery(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the newly inserted album as JSON
	c.JSON(http.StatusCreated, album)
}

// GetAlbum returns a single album from the database.
func GetAlbum(c *gin.Context) {
	// Get the ID parameter from the URL
	id := c.Param("id")
	// Call the GetAlbumQuery function to get the album
	album, err := queries.GetAlbumQuery(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting album"})
		return
	}
	// Return the album as JSON
	c.JSON(http.StatusOK, album)
}


// DeleteAlbum deletes an album from the database.
func DeleteAlbum(c *gin.Context) {
	// Get the ID parameter from the URL
	id := c.Param("id")
	// Call the DeleteAlbumQuery function to delete the album
	err := queries.DeleteAlbumQuery(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error deleting album"})
		return
	}
	// Return a 204 status code
	c.Status(http.StatusNoContent)
}

// PatchAlbum updates an album in the database.
func PatchAlbums(c *gin.Context) {
	// Get the ID parameter from the URL
	id := c.Param("id")
	// Validate input
	var req models.Album
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Call the PutAlbumQuery function to update the album
	album, err := queries.PatchAlbumQuery(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error updating album"})
		return
	}
	// Return the updated album as JSON
	c.JSON(http.StatusOK, album)
}


