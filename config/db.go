package config

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Driver   string
	Username string
	Password string
	Name     string
	Host     string
	Port     string
	Timezone string
	SslMode  bool
}

func (db *DatabaseConfig) GetDSN() string {
	db.Host = os.Getenv("DB_HOST")
	db.Username = os.Getenv("POSTGRES_USER")
	db.Password = os.Getenv("POSTGRES_PASSWORD")
	db.Name = os.Getenv("POSTGRES_NAME")
	db.Port = os.Getenv("DB_PORT")
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Baghdad",
		db.Host, db.Username, db.Password, db.Name, db.Port)
}
