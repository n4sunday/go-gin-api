package controllers

import (
	models "go-gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindFirearms(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var firearms []models.Firearms
	db.Find(&firearms)
	c.JSON(http.StatusOK, gin.H{"data": firearms})
}

func CreateFirearms(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input models.CreateFirearmsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	firearms := models.Firearms{
		Name:     input.Name,
		Year:     input.Year,
		Caliber:  input.Caliber,
		System:   input.System,
		Capacity: input.Capacity,
		Barrel:   input.Barrel,
		Size:     input.Size,
	}
	db.Create(&firearms)
	c.JSON(http.StatusCreated, gin.H{"data": firearms})
}
