package main

import (
	// "log"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/user/:id/type/:type", GetUserById)
	router.GET("/product", productHandler)

	router.POST("/books", storeBooks)

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}

func helloHandler(c *gin.Context) {
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

func productHandler(c *gin.Context) {
	product := c.Query("name")
	price := c.Query("price")
	c.JSON(200, gin.H{
		"message": "Product " + product + " Price " + price,
	})

}

type Book struct {
	Name     string `json:"name" binding:"required"`
	Price    json.Number    `json:"price" binding:"required,number"`
	SubTitle string `json:"sub_title"`
}

func storeBooks(c *gin.Context) {
	var book Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errorMessage,
			})
			return
		}

	}

	c.JSON(200, gin.H{
		"name":      book.Name,
		"price":     book.Price,
		"sub_title": book.SubTitle,
	})
}
