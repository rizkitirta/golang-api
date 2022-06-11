package handler

import (
	"fmt"
	"golang-api-gin/book"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (handler *bookHandler) RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}

func (handler *bookHandler) HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello golang",
	})
}

func (handler *bookHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	tipe := c.Param("type")
	c.JSON(200, gin.H{
		"message": "hello " + id + " Tipe " + tipe,
	})
}

func (handler *bookHandler) ProductHandler(c *gin.Context) {
	product := c.Query("name")
	price := c.Query("price")
	c.JSON(200, gin.H{
		"message": "Product " + product + " Price " + price,
	})

}

func (handler *bookHandler) StoreBooks(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := handler.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{
		"data": book,
	})
}

func (handler *bookHandler) StoreBooksV2(c *gin.Context) {
	var book book.BookRequestV2

	err := c.ShouldBindJSON(&book)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(200, gin.H{
		"name":      book.Title,
		"price":     book.Price,
		"sub_title": book.SubTitle,
	})
}
