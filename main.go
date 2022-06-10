package main

import (
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
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
	Price    int    `json:"price" binding:"required,number"`
	SubTitle string `json:"sub_title"`
}

func storeBooks(c *gin.Context) {
	var book Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"name":      book.Name,
		"price":     book.Price,
		"sub_title": book.SubTitle,
	})
}
