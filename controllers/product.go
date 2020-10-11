package controllers

import (
	models "go-gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindItem(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var item []models.Item
	db.Find(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func CreateItem(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Validate input
	var input models.CreateItem
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create Book
	item := models.Item{ItemName: input.ItemName, Amount: input.Amount}
	db.Create(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func FindOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var order []models.Order
	db.Find(&order)
	c.JSON(http.StatusOK, gin.H{"data": order})
}

func FindOrderItem(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var orderItem []models.OrderItem
	db.Find(&orderItem)
	c.JSON(http.StatusOK, gin.H{"data": orderItem})
}

func CreateOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	order := models.Order{Status: "pending"}
	db.Create(&order)
	item := models.OrderItem{OrderID: order.ID, ItemID: 2, Quantity: 4}
	db.Create(&item)
	c.JSON(http.StatusOK, gin.H{"data": order})
}
