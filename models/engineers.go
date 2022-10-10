package models

import (
	"github.com/jeki-aka-zer0/go-crud/config"

	"gorm.io/gorm"
)

type Engineer struct {
	ID        int64
	FirstName string
	LastName  string
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetAllEngineers() []Engineer {
	var engineers []Engineer
	db.Find(&engineers)
	return engineers
}

func GetEngineerById(Id int64) *Engineer {
	var engineer Engineer
	db.Where("ID = ?", Id).Find(&engineer)
	return &engineer
}

func (e *Engineer) CreateEngineer() *Engineer {
	db.Create(&e)
	return e
}
