package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kd-o-guh/Library-API/models"
)

var books = []models.Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func createBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := bookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)

}

func bookByID(id string) (*models.Book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}

	book, err := bookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not avaliable"})
		return
	}

	book.Quantity -= 1

	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}

	book, err := bookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func RouterInit() *gin.Engine {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	return router
}
