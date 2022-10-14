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

func GetEngineerById(ID int64) *Engineer {
	var engineer Engineer
	db.Where("ID = ?", ID).Find(&engineer)
	return &engineer
}

func (e *Engineer) CreateEngineer() *Engineer {
	db.Create(&e)
	return e
}

func DeleteEngineer(ID int64) Engineer {
	var engineer Engineer
	db.Where("ID = ?", ID).Delete(engineer)
	return engineer
}

func UpdateEngineer(ID int64, engineerData Engineer) *Engineer {
	existentEngineer := GetEngineerById(ID)

	if engineerData.FirstName != "" {
		existentEngineer.FirstName = engineerData.FirstName
	}
	if engineerData.LastName != "" {
		existentEngineer.LastName = engineerData.LastName
	}

	db.Save(&existentEngineer)
	return existentEngineer
}
