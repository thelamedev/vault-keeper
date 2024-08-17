package database

import (
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/sqlite"

	"github.com/thelamedev/vault-keeper/internal/config"
)

var DB *gorm.DB

func Connect(config *config.Config) error {
	var err error
	DB, err = gorm.Open(config.Database.Driver, config.Database.Filename)
	return err
}
