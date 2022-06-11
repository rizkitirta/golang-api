package main

import (
	"golang-api-gin/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/golang_api?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

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
