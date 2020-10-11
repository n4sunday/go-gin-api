package models

type User struct {
	ID       uint   `json:"id" gorm:"PrimaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Last     string `json:"last"`
}

type CreateUser struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Last     string `json:"last" binding:"required"`
}
