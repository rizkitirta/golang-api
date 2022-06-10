package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)

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