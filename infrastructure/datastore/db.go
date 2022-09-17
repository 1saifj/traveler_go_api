package datastore

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"taveler/infrastructure/model"
)

func SetupDB() (*gorm.DB, error) {
	var dialector gorm.Dialector
	//dbConfig := config.DatabaseConfig{}
	//dialector = postgres.Open(dbConfig.GetDSN())
	dialector = sqlite.Open("taveler.db")
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = migrate(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(model.Place{}, model.Category{}, model.Image{}, model.User{})
}
