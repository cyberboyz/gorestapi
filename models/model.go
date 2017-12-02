package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

var db *gorm.DB
var err error

func init() {
	db_url := os.Getenv("DATABASE_URL")
	if db_url == "" {
		db_url = "host=localhost user=postgres dbname=gorm sslmode=disable password=postgres"
	}
	db, err = gorm.Open("postgres", db_url)
	db.SingularTable(true)
	if err != nil {
		log.Panic(err)
	}
	db.AutoMigrate(User{})
}
