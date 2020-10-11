package main

import (
	"fmt"
	controllers "go-gin-api/controllers"
	models "go-gin-api/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Get ENV
	var errEnv error
	errEnv = godotenv.Load()
	if errEnv != nil {
		log.Fatalf("Error getting env, %v", errEnv)
	}
	host := os.Getenv("POSTGRES_HOST")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	database := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", host, user, password, database, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(
		&models.Book{},
		&models.Firearms{},
		&models.Brand{},
		&models.Item{},
		&models.Order{},
		&models.OrderItem{},
		&models.Todo{},
		&models.User{},
	)

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Go Gin",
		})
	})
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("/firearms", controllers.FindFirearms)
	r.POST("/firearms", controllers.CreateFirearms)

	r.GET("/item", controllers.FindItem)
	r.POST("/item", controllers.CreateItem)
	r.GET("/order", controllers.FindOrder)
	r.POST("/order", controllers.CreateOrder)
	r.GET("/orderitem", controllers.FindOrderItem)

	todo := r.Group("/todo")
	{
		todo.GET("/", controllers.FindTodo)
		todo.POST("/", controllers.CreateTodo)
		todo.GET("/:userid", controllers.FindTodoByUserID)
	}

	r.POST("/upload", controllers.Upload)
	r.StaticFS("/file", http.Dir("public"))

	r.GET("/user/:id", controllers.FindUserById)
	r.GET("/user", controllers.FindUser)
	r.POST("/user", controllers.CreateUser)

	r.Run(":9500")
}
