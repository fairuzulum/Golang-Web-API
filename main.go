package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}

func inputBookHandler(c *gin.Context) {
	var input InputBook

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"title": input.Title,
		"price": input.Price,
	})
}
