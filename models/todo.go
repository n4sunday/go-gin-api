package models

type Todo struct {
	ID      uint   `json:"id" gorm:"PrimaryKey"`
	UserID  int    `gorm:"ForeignKey:userID"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type CreateTodo struct {
	UserID  int    `json:"userID" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Message string `json:"message" binding:"required"`
}
