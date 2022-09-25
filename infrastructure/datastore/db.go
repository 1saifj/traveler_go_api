package datastore

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"taveler/infrastructure/model"
)

var (
	DB *gorm.DB
)

func SetupDB() (*gorm.DB, error) {
	var dialector gorm.Dialector
	//dbConfig := config.DatabaseConfig{}
	//dialector = postgres.Open(dbConfig.GetDSN())
	dialector = sqlite.Open("taveler.db")
	db, err := gorm.Open(dialector, &gorm.Config{})
	//permify, _ := permifry.New(permifry.Options{
	//	Migrate: true,
	//	DB:      db,
	//})
	//err = permify.CreateRole("admin", "admin")
	//err = permify.CreatePermission("admin", "control all data")

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
	return db.AutoMigrate(model.File{}, model.Place{}, model.Category{}, model.Image{}, model.User{})
}
