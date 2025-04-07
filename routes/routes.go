package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/trae/go-restful-crud-api/controllers"
)

// SetupRouter configures the API routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Group routes with the /api prefix
	api := router.Group("/api")
	{
		// Book routes
		api.GET("/books", controllers.GetBooks)
		api.GET("/books/:id", controllers.GetBook)
		api.POST("/books", controllers.CreateBook)
		api.PUT("/books/:id", controllers.UpdateBook)
		api.DELETE("/books/:id", controllers.DeleteBook)
	}

	return router
}