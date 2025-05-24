package main

import (
	"example/db"
	"net/http"
	"example/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDB()
	router := gin.Default()
	router.LoadHTMLGlob("view/*")
	router.GET("/view", func(c *gin.Context) {
		c.HTML(http.StatusOK, "view.html",gin.H{})
	})
	router.POST("/add",handler.AddBookHandler)
	router.GET("/books", handler.ViewBooks)
	router.POST("/update",handler.UpdateBookHandler)
	router.POST("/delete",handler.DeleteBookHandler)
	router.Run(":8083")

}
