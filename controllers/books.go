package controllers

import (
	"net/http"
	"web-service-gin/dtos"
	"web-service-gin/models"
	"web-service-gin/services"

	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)

	c.IndentedJSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	var bookDto dtos.CreateBookDto

	if err := c.ShouldBindJSON(&bookDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	book := services.CreateBook(bookDto)

	models.DB.Create(&book)

	c.IndentedJSON(http.StatusCreated, gin.H{"data": book})
}

func FindBookById(c *gin.Context) {
	book, err := services.GetBookById(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	var bookDto dtos.UpdateBookDto

	if err := c.ShouldBindJSON(&bookDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := services.UpdateBook(c.Param("id"), bookDto)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "record not found"})
	}

	models.DB.Model(&book).Updates(bookDto)

	c.IndentedJSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	err := services.DeleteBook(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{})
}
