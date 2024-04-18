package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var defaultDB *gorm.DB

func PersistDB() *gorm.DB {
	return defaultDB
}

func PersistConnect() error {
	if defaultDB != nil {
		return nil
	}

	con, err := NewConnection()

	if err != nil {
		return err
	}

	defaultDB = con

	return nil
}

func DoPersistConnect() error {
	return PersistConnect()
}

func NewConnection() (*gorm.DB, error) {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	password := viper.GetString("database.password")
	user := viper.GetString("database.user")
	dbname := viper.GetString("database.dbname")

	dsn := fmt.Sprintf("host=%s port=%s password=%s user=%s dbname=%s sslmode=disable",
		host, port, password, user, dbname)

	con, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	return con, nil
}
