package handler

import (
	"fmt"
	"golang-api-gin/book"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
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

	result, err := handler.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": book.ResponseConverter(result),
	})
}

func (handler *bookHandler) UpdateBook(c *gin.Context) {
	var bookUpdateRequest book.BookUpdateRequest

	err := c.ShouldBindJSON(&bookUpdateRequest)
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

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	result, err := handler.bookService.Update(id, bookUpdateRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": book.ResponseConverter(result),
	})
}

func (handler *bookHandler) GetBooks(c *gin.Context) {
	books, err := handler.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var booksResponse []book.BookResponse
	for _, v := range books {
		bookResponse := book.ResponseConverter(v)

		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (handler *bookHandler) GetBookById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.bookService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookResponse := book.ResponseConverter(result)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (handler *bookHandler) Delete(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.bookService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookResponse := book.ResponseConverter(result)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}
