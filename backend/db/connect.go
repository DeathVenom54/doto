package db

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() error {
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_CONN_STRING")), &gorm.Config{})
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&User{})
	if err != nil {
		return err
	}

	return nil
}
