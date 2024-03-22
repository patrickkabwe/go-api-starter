package db

import (
	"github.com/patrickkabwe/go-api-starter/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.Env.DatabaseUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
