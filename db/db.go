package db

import (
	"github.com/patrickkabwe/go-api-starter/config"
	"github.com/patrickkabwe/go-api-starter/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.Env.DatabaseUrl), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	if err = Migrate(db); err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	return db.AutoMigrate(&types.User{})
}
