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

	// Create
	// book := book.Book{
	// 	Title: "Belajar Golang",
	// 	Price: 100,
	// 	Description: "Mahir menggunakan golang",
	// 	Rating: 5,
	// }
	// db.Create(&book)
	
	// Find One
	// var book book.Book
	// err = db.Debug().First(&book,1).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Title: ", book.Title, " Price: ", book.Price, " Description: ", book.Description, " Rating: ", book.Rating)

	// Find All
	// var books []book.Book
	// err = db.Debug().Where("rating = ?",5).Find(&books).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, book := range books {
	// 	log.Println("Title: ", book.Title, " Price: ", book.Price, " Description: ", book.Description, " Rating: ", book.Rating)
	// }

	// Update
	var book book.Book
	err = db.Debug().Where("id= ?",3).First(&book).Error
	if err != nil {
		log.Fatal(err)
	}
	book.Title = "Belajar Golang (Update)"
	err = db.Debug().Save(&book).Error
	if err != nil {
		log.Fatal(err)
	}

	// Router
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
