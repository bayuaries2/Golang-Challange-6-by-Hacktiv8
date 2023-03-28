package routers

import (
	"Challanges-6/controllers"

	"github.com/gin-gonic/gin"
)

func StarServer() *gin.Engine {
	router := gin.Default()

	router.POST("/book", controllers.CreateBook)

	router.PUT("/book/:id", controllers.UpdateBookById)

	router.GET("/books", controllers.GetAllBook)

	router.GET("/book/:id", controllers.GetBookById)

	router.DELETE("/book/:id", controllers.DeleteBookById)

	return router
}
