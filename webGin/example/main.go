package main

import (
	"test/controllers"
	"github.com/gin-gonic/gin"

)

func main() {

	// Set the router as the default
	r := gin.Default()


	// Migrate the schema
	// ----------------------------------------------------------------
	// dbConnection.Migrate()

	// Group routes
	// ----------------------------------------------------------------
	// * Group routes by version
	v1 := r.Group("/v1")
	{
		// * GET a list of albums
		v1.GET("/albums", controllers.GetAlbums)

		// * POST a new album
		v1.POST("/albums", controllers.PostAlbums)
		
		// * GET a single album
		v1.GET("/albums/:id", controllers.GetAlbum)

		// * DELETE to remove a single album
		v1.DELETE("/albums/:id", controllers.DeleteAlbum)

		// * Patch to update a single album
		v1.PATCH("/albums/:id", controllers.PatchAlbums)

	}


	// * Listen and serve on
	r.Run(":8080")
}
