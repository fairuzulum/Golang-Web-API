package handler

import (
	"fmt"
	"golang-web-api/book"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func BookHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"Book ID ": id,
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"book Id ":    id,
		"Book Title ": title,
	})
}

func InputBookHandler(c *gin.Context) {
	var input book.InputBook

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
