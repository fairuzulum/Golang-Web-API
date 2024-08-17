package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/book/:id", bookHandler)
	router.GET("/books", queryHandler) 

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func bookHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"Book ID ": id,
	})
}

func queryHandler(c *gin.Context){
	title := c.Query("title")
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"book Id ": id,
		"Book Title ": title,
	})
}
