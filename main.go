package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/book/:id", bookHandler)
	router.GET("/query", queryHandler)

	router.POST("/books", inputBookHandler)

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

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"book Id ":    id,
		"Book Title ": title,
	})
}

type InputBook struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required,number"`
}

func inputBookHandler(c *gin.Context) {
	var input InputBook

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": input.Title,
		"price": input.Price,
	})
}
