package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id", booksHandler)
	router.GET("/movies", moviesHandler)
	router.Run()
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name":    "Fuad Muhammad Nur",
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
