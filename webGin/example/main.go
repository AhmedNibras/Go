package main

import (
	"test/controllers"

	"github.com/gin-gonic/gin"
)



func main() {
	
	r := gin.Default()

	r.GET("/albums", controllers.GetAlbums)
	r.GET("/albums/:id", controllers.GetAlbumsByID)
	r.POST("/albums", controllers.PostAlbums)
	
	r.Run(":8080")
}