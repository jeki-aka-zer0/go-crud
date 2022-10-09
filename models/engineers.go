package models

import (
	"github.com/jeki-aka-zer0/go-crud/config"

	"gorm.io/gorm"
)

type Engineer struct {
	ID        int
	FirstName string
	LastName  string
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func Test() {
	// Create
	db.Create(&Engineer{FirstName: "John", LastName: "Doe"})
}
