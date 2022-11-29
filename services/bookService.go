package services

import (
	"web-service-gin/dtos"
	"web-service-gin/models"
	"web-service-gin/repository"
)

func GetBooks() []models.Book {
	return repository.GetBooks()
}

func CreateBook(bookDto dtos.CreateBookDto) models.Book {
	return repository.CreateBook(bookDto)
}

func GetBookById(id string) (models.Book, error) {
	book, err := repository.FindBookById(id)

	return book, err
}

func UpdateBook(id string, bookDto dtos.UpdateBookDto) (models.Book, error) {
	return repository.UpdateBook(id, bookDto)
}

func DeleteBook(id string) error {
	return repository.DeleteBook(id)
}
