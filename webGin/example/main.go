package main

import (
	"test/controllers"
	"test/dbConnection"
	"github.com/gin-gonic/gin"
)

func main() {

	// Set the router as the default
	r := gin.Default()


	// Connect to the database
	// ----------------------------------------------------------------
	dbConnection.Connection()

	// Migrate the schema
	// ----------------------------------------------------------------
	// dbConnection.Migrate()

	// Routes
	// ----------------------------------------------------------------
	// GET a list of albums
	r.GET("/albums", controllers.GetAlbums)

	// GET a single album
	r.GET("/albums/:id", controllers.GetAlbumsByID)

	// POST a new album
	r.POST("/albums", controllers.PostAlbums)

	// PUT to update a new album
	r.PUT("/albums/:id", controllers.PutAlbums)

	// DELETE an album
	r.DELETE("/albums/:id", controllers.DeleteAlbums)

	// PATCH to update an album's artist
	r.PATCH("/albums/:id", controllers.PatchAlbums)

	// Listen and serve on
	r.Run(":8080")
}
