package main

import (
	"web-service-gin/controllers"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBookById)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	models.ConnectDatabase()

	r.Run("localhost:8080")
}
