package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func (db *db) handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": db.data,
	})
}

type db struct {
	data []string
}

type DataRequest struct {
	Text string `json:"text"`
}

func (db *db) postHandler(c *gin.Context) {
	var data DataRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.data = append(db.data, data.Text)
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dikirim", "data": data.Text})
}

func main() {
	database := db{}
	database.data = make([]string, 0)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))
	r.GET("/", database.handler)
	r.POST("/send", database.postHandler)
	r.Run(":" + os.Getenv("PORT"))
}
