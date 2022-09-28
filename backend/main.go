package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": db,
	})
}

type database struct {
	data []string
}

var db []string

type DataRequest struct {
	Text string `json:"text"`
}

func postHandler(c *gin.Context) {
	var data DataRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db = append(db, data.Text)
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dikirim", "data": data.Text})
}

func main() {
	db = make([]string, 0)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))
	r.GET("/", handler)
	r.POST("/send", postHandler)
	r.Run(":" + os.Getenv("PORT"))
}
