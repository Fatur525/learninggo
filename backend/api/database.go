package api

import (
	"os"

	"github.com/Fatur525/learninggo/model"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDb() (*gorm.DB, error) {
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

	if err := dbconn.AutoMigrate(&model.Todos{}); err != nil {
		panic("Failed to migrate database")
	}

	return dbconn, err
}
