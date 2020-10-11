package controllers

import (
	"fmt"
	"go-gin-api/models"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var todo []models.Todo
	db.Find(&todo)
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func CreateTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input models.CreateTodo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{
		Username: input.Username,
		Title:    input.Title,
		Message:  input.Message,
	}
	db.Create(&todo)
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("public/" + filename)
	if err != nil {
		// log.Fatal{err}
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		// log.Fatal{err}
	}
	filepath := "http://localhost:9500/file/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}
