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
	a, err := queries.GetAlbumsQuery()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting albums"})
		return
	}
	// Return the list of albums as JSON
	c.JSON(http.StatusOK, a)
}

// PostAlbums adds an array of album to the database.
func PostAlbums(c *gin.Context) {
	// Validate input
	var m []models.Album
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Call the PostAlbumQuery function to add the album
	a, err := queries.PostAlbumQuery(m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error adding album"})
		return
	}
	// Return the newly added album as JSON
	c.JSON(http.StatusCreated, a)
}

// GetAlbum returns a single album from the database.
func GetAlbum(c *gin.Context) {
	// Get the ID parameter from the URL
	id := c.Param("id")
	// Call the GetAlbumQuery function to get the album
	a, err := queries.GetAlbumQuery(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Album Not Found"})
		return
	}
	// Return the album as JSON
	c.JSON(http.StatusOK, a)
}

// DeleteAlbum deletes an album from the database.
func DeleteAlbum(c *gin.Context) {
	// Get the ID parameter from the URL
	id := c.Param("id")
	

	//Check if the album is already in the database
	if _, err := queries.GetAlbumQuery(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	// Call the DeleteAlbumQuery function to delete the album
	if err := queries.DeleteAlbumQuery(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error deleting album"})
		return
	}
	
	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "album deleted"})
}

// PatchAlbum updates an album in the database.
func PatchAlbum(c *gin.Context) {
	// Get the ID parameter from the URL
	id := c.Param("id")

	// If the album already exists
	if _, err := queries.GetAlbumQuery(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	// Validate input
	var m models.Album
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Call the PatchAlbumQuery function to update the album
	a, err := queries.PatchAlbumQuery(id, &m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error updating album"})
		return
	}
	// Return the album as JSON
	c.JSON(http.StatusOK, a)
}
