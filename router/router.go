package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize(){
	// router definition 
	router := gin.Default()
	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hello from root",
		})
	})
	// run API, access on "https://0.0.0.0:8080"
	router.Run("0.0.0.0:8080")
}