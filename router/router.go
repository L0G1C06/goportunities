package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize(){
	// Initialize Router 
	router := gin.Default()

	// Initialize Routes
	InitializeRoutes(router)
	
	// run the server, access on "https://0.0.0.0:8080"
	router.Run("0.0.0.0:8080")
}