package main

import (
	"golang-api-gin/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	API_V1 := router.Group("/api/v1")
	API_V1.GET("/", handler.RootHandler)
	API_V1.GET("/hello", handler.HelloHandler)
	API_V1.GET("/user/:id/type/:type", handler.GetUserById)
	API_V1.GET("/product", handler.ProductHandler)
	API_V1.POST("/books", handler.StoreBooks)

	API_V2 := router.Group("/api/v2")
	API_V2.POST("/books", handler.StoreBooksV2)

	router.Run()
}




