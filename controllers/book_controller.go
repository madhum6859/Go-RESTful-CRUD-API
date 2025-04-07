package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trae/go-restful-crud-api/models"
)

// books slice to simulate a database
var books = []models.Book{
	{ID: "1", Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Year: 2015},
	{ID: "2", Title: "Clean Code", Author: "Robert C. Martin", Year: 2008},
	{ID: "3", Title: "Design Patterns", Author: "Erich Gamma", Year: 1994},
}

// GetBooks returns all books
func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

// GetBook returns a specific book by ID
func GetBook(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

// CreateBook creates a new book
func CreateBook(c *gin.Context) {
	var newBook models.Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

// UpdateBook updates an existing book
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.Book

	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, book := range books {
		if book.ID == id {
			updatedBook.ID = id
			books[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

// DeleteBook deletes a book
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}