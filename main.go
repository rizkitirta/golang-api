package main

import (
	"golang-api-gin/book"
	"golang-api-gin/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/golang_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// migrate the schema
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// Router
	router := gin.Default()

	API_V1 := router.Group("/api/v1")
	API_V1.GET("/", bookHandler.RootHandler)
	API_V1.GET("/hello", bookHandler.HelloHandler)
	API_V1.GET("/user/:id/type/:type", bookHandler.GetUserById)
	API_V1.GET("/product", bookHandler.ProductHandler)
	API_V1.POST("/books", bookHandler.StoreBooks)

	API_V2 := router.Group("/api/v2")
	API_V2.POST("/books", bookHandler.StoreBooksV2)

	router.Run()
}
