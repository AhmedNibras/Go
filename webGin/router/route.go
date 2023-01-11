package router

import (
	"test/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the router
func SetupRouter() *gin.Engine {

	// Set the router as the default
	r := gin.Default()
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
		v1.PATCH("/albums/:id", controllers.PatchAlbum)

	}

	// * Listen and serve on
	r.Run(":8080")
	return r
}