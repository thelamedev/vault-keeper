package entities

import (
	"github.com/jinzhu/gorm"
)

type APIKey struct {
	gorm.Model

	Key     string `gorm:"unique;not null"`
	UserId  uint   `gorm:"index"`
	Revoked bool
}
