package repository

import (
	"web-service-gin/dtos"
	"web-service-gin/models"
)

func GetBooks() []models.Book {
	var books []models.Book

	models.DB.Find(&books)

	return books
}

func CreateBook(bookDto dtos.CreateBookDto) models.Book {

	book := models.Book{Title: bookDto.Title, Author: bookDto.Author}

	models.DB.Create(&book)

	return book
}

func FindBookById(id string) (models.Book, error) {
	var book models.Book

	if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func UpdateBook(id string, bookDto dtos.UpdateBookDto) (models.Book, error) {
	book, err := FindBookById(id)

	if err != nil {
		return book, err
	}

	models.DB.Model(&book).Updates(bookDto)

	return book, nil
}

func DeleteBook(id string) error {
	book, err := FindBookById(id)

	if err != nil {
		return err
	}

	models.DB.Delete(book)

	return nil
}
