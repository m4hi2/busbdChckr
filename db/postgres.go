package db

import (
	"fmt"
	"github.com/m4hi2/busbdChckr/db/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	Db  *gorm.DB
	err error
)

func ConnectDB() *gorm.DB {
	if Db != nil {
		return Db
	}
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	password := viper.GetString("database.password")
	user := viper.GetString("database.user")
	dbname := viper.GetString("database.dbname")

	dsn := fmt.Sprintf("host=%s port=%s password=%s user=%s dbname=%s sslmode=disable",
		host, port, password, user, dbname)

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error", err)
		return nil
	}
	err = Db.AutoMigrate(&models.User{}, &models.Log{})
	if err != nil {
		log.Fatal("DB migration error: ", err)
		return nil
	}
	fmt.Println("Successfully Connected Database")
	return Db
}