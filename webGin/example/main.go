package main

import (
	"test/router"
	"github.com/gin-gonic/gin"

)

func main() {
	// Set the gin mode
	// ----------------------------------------------------------------
	gin.SetMode(gin.DebugMode)

	// Migrate the schema
	// ----------------------------------------------------------------
	// dbConnection.Migrate()

	// Setup the router
	// ----------------------------------------------------------------
	router.SetupRouter()

}
