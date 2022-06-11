package handler

import (
	"fmt"
	"golang-api-gin/book"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello golang",
	})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	tipe := c.Param("type")
	c.JSON(200, gin.H{
		"message": "hello " + id + " Tipe " + tipe,
	})
}

func ProductHandler(c *gin.Context) {
	product := c.Query("name")
	price := c.Query("price")
	c.JSON(200, gin.H{
		"message": "Product " + product + " Price " + price,
	})

}


func StoreBooks(c *gin.Context) {
	var book book.BookInput
	
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
		"name":      book.Name,
		"price":     book.Price,
		"sub_title": book.SubTitle,
	})
}

func StoreBooksV2(c *gin.Context) {
	var book book.BookInputV2
	
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
