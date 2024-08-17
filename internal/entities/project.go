package entities

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model

	Name        string `gorm:"unique;not null"`
	Description string
	OwnerId     uint `gorm:"index"`
}
