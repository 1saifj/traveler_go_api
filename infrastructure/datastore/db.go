package datastore

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
	"taveler/config"
	"taveler/infrastructure/model"
)

func SetupDB(con *config.DatabaseConfig) (*gorm.DB, error) {
	var dialector gorm.Dialector
	switch strings.ToLower(con.Driver) {
	case "postgresql":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%v TimeZone=%s",
			con.Host, con.Username, con.Password, con.Name, con.Port, con.SslMode, con.Timezone)
		dialector = postgres.Open(dsn)
		break
	case "sqlite":
		if !strings.HasSuffix(con.Name, ".db") {
			con.Name += ".db"
		}
		dialector = sqlite.Open(con.Name)
		break
	default:
		return nil, errors.New("invalid DB driver")
	}

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
	return db.AutoMigrate(model.Place{}, model.Category{})
}
