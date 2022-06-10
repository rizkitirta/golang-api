package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/user/:id/type/:type", GetUserById)
	router.GET("/product",productHandler)

	router.POST("/books",storeBooks)

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}

func helloHandler(c *gin.Context)  {
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
	Name string  
	Price int  
}

func storeBooks(c *gin.Context)  {
	var book Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"name": book.Name,
		"price": book.Price,
	})
}
	