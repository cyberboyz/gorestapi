package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type Response struct {
	Message []string    `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var db *gorm.DB
var err error

func init(){

	db, err = gorm.Open("postgres", "postgres://yvqfzerfajfbkw:c6d2eefb105837e44465c0dfe2cd55a68a8621ee6c3bc9a52ecf57f1c934b00e@ec2-54-163-246-154.compute-1.amazonaws.com:5432/de64t12g30v1rj")
	db.SingularTable(true)
	if err != nil {
		log.Panic(err)
	}

}