package main

import (
	"golang-api-gin/book"
	"golang-api-gin/handler"
	"golang-api-gin/user"
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
	db.AutoMigrate(&user.User{})

	//book handlers
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	//user handlers
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Router
	router := gin.Default()

	API_V1 := router.Group("/api/v1")
	API_V1.POST("/books", bookHandler.StoreBooks)
	API_V1.PUT("/books/:id", bookHandler.UpdateBook)
	API_V1.GET("/books", bookHandler.GetBooks)
	API_V1.GET("/book/:id", bookHandler.GetBookById)
	API_V1.DELETE("/book/:id", bookHandler.Delete)

	// User API
	API_V1.GET("/users", userHandler.GetUser)
	API_V1.POST("/user", userHandler.Store)
	API_V1.PUT("/user", userHandler.UpdateUser)
	API_V1.DELETE("/user/:id", userHandler.Delete)
	router.Run()
}
