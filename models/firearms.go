package models

import (
	"time"

	"gorm.io/gorm"
)

type Firearms struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Year      int    `json:"year"`
	Caliber   string `json:"caliber"`
	System    string `json:"system"`
	Capacity  string `json:"capacity"`
	Barrel    string `json:"barrel"`
	Size      string `json:"size"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Brand struct {
	gorm.Model
	BrandID  uint       `json:"brandID" gorm:"primary_key"`
	Name     string     `json:"name"`
	Country  string     `json:"country"`
	firearms []Firearms `gorm:"many2many:firearms"`
}
