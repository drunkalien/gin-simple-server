package controllers

import (
	"net/http"
	"web-service-gin/dtos"
	"web-service-gin/models"

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

	book := models.Book{Title: bookDto.Title, Author: bookDto.Author}
	models.DB.Create(&book)

	c.IndentedJSON(http.StatusCreated, gin.H{"data": book})
}

func FindBookById(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	var bookDto dtos.UpdateBookDto

	if err := c.ShouldBindJSON(&bookDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(bookDto)

	c.IndentedJSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	models.DB.Delete(book)

	c.IndentedJSON(http.StatusOK, gin.H{})
}
