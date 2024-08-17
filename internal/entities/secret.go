package entities

import (
	"github.com/jinzhu/gorm"
)

type Secret struct {
	gorm.Model

	Name      string `gorm:"unique;not null"`
	Value     string `gorm:"not null"`
	ProjectId uint   `gorm:"index"`
}
