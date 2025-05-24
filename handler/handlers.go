package handler

import (
	"example/controls"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddBookHandler(c *gin.Context) {
	var req struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "JSON error"})
		return
	}

	err := controls.InsertBook(req.Title, req.Author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Error in inserting "})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Messsage": "Inserted"})
}

func DeleteBookHandler(c *gin.Context) {
	var req struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("JSON Binding Error:", err) 
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSON problem"})
		return
	}

	err := controls.DeleteBook(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Deleting error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Deleted"})
}

func UpdateBookHandler(c *gin.Context) {
	var req struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("JSON Binding Error:", err) 
		c.JSON(http.StatusBadRequest, gin.H{"Message": "JSON problem"})
		return
	}

	err := controls.UpdateBook(req.ID, req.Title, req.Author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Update problem"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Updated"})
}

func ViewBooks(c * gin.Context){
	books, err := controls.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching books"})
		return
	}
	c.JSON(http.StatusOK, books)
}
