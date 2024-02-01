package router

import (
	"github.com/L0G1C06/goportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine){
	// Initialize handler (logger and database)
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}