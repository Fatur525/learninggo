package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataRequest struct {
	Text string `json:"text" binding:"required"`
}

type Todos struct {
	gorm.Model
	Task string `json:"task"`
	Done bool
}

type repository struct {
	db *gorm.DB
}

func (r *repository) handler(c *gin.Context) {
	var todos []Todos
	res := r.db.Find(&todos)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": todos,
	})
}

func (r *repository) postHandler(c *gin.Context) {
	var data DataRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := Todos{
		Task: data.Text,
		Done: false,
	}

	res := r.db.Create(&todo)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dikirim", "data": todo})
}

func main() {
	var dbconn *gorm.DB
	var err error
	dbUrl := os.Getenv("DATABASE_URL")
	if os.Getenv("ENVIRONMENT") == "PROD" {
		dbconn, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	} else {
		dbconn, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	}

	if err != nil {
		panic("Failed to connect database")
	}

	sqlDB, err := dbconn.DB()
	if err != nil {
		panic("Failed to get database")
	}

	if err := sqlDB.Ping(); err != nil {
		panic("Failed to get database")
	}

	if err := dbconn.AutoMigrate(&Todos{}); err != nil {
		panic("Failed to migrate database")
	}

	repo := repository{
		db: dbconn,
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))
	r.GET("/", repo.handler)
	r.POST("/send", repo.postHandler)

	r.Run(":" + os.Getenv("PORT"))
}
