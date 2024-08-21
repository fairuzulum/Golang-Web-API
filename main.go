package main

import (
	"golang-web-api/book"
	"golang-web-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//main
	//handler
	//service
	//repository
	//db

	dsn := "root:@tcp(127.0.0.1:3306)/golang_web_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// book, _ := bookService.FindAll()
	// fmt.Println(book)

	// book, _ := bookService.FindByID(3)
	// fmt.Println(book)

	newBook := book.BookRequest{
		Title: "Cinta monyet",
		Price: "50000",
	}
	bookService.Create(newBook)

	/////////////////////////////////////////////

	// books, err := bookRepository.FindAll()
	// for _, book := range books {
	// 	fmt.Println("Title: ", book.Title)
	// }

	// newBook := book.Book{
	// 	Title:       "Gak tau",
	// 	Description: "KApa aja Udah yang penting gak tau",
	// 	Price:       50000,
	// 	Rating:      9,
	// 	Discount:    70,
	// }
	// bookRepository.Create(newBook)

	// book, err := bookRepository.FindByID(3)
	// fmt.Println(book)

	///////////////////////////////////

	// book := book.Book{
	// 	Title:       "Gak tau",
	// 	Description: "KApa aja Udah yang penting gak tau",
	// 	Price:       50000,
	// 	Rating:      9,
	// 	Discount:    70,
	// }
	// // CREATE
	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("CREATE DATA FAILED")
	// }

	// var book book.Book

	// // FIND BY ID
	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("DATA NOT FOUND")
	// }

	// fmt.Println("title:" + book.Title)
	// fmt.Println(book)

	// UPDATE
	// book.Title = "uuuuuuuuuuuu"
	// err = db.Debug().Save(&book).Error
	// if err != nil {
	// 	fmt.Println("UPDATE DATA FAILED")
	// }

	// // DELETE
	// err = db.Debug().Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("DELETE DATA FAILED")
	// }

	router := gin.Default()

	v1 := router.Group("/api/v1")

	// v1.GET("/", bookHandler.RootHandler)
	// v1.GET("/book/:id", bookHandler.BooksHandler)
	v1.GET("/books", bookHandler.GetAllBooksHandler)
	v1.GET("/books/:id", bookHandler.GetBookByIDHandler)
	v1.POST("/books", bookHandler.InputBookHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	router.Run()
}
