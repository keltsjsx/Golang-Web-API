package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id", booksHandler)
	router.GET("/movies", moviesHandler)
	router.POST("/books", postBooksHandler)
	router.Run()
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name":    "Keltskaya",
		"address": "Depok",
	})
}

func helloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title":    "Hello World",
		"subtitle": "Belajar golang Web API dari dasar banget",
	})
}

func booksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func moviesHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	ctx.JSON(http.StatusOK, gin.H{"title": title})
}

type BookInput struct {
	Title    string
	Price    int
	SubTitle string `json:"sub_title"`
}

func postBooksHandler(ctx *gin.Context) {
	var bookInput BookInput

	err := ctx.ShouldBindJSON(&bookInput)

	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})
}
