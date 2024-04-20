package repos

import (
	"github.com/m4hi2/busbdChckr/db/models"
	"gorm.io/gorm"
	"log"
)

type UserStore struct {
	DB *gorm.DB
}

func (d *UserStore) GetAll() ([]models.User, error) {
	var users []models.User

	err := d.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, err
}

func (d *UserStore) FindHash(hash string) (*models.Log, error) {
	log := &models.Log{}
	res := d.DB.Where("hash = ?", hash).First(log)
	if res.Error != nil {
		return nil, res.Error
	}
	return log, nil
}

func (d *UserStore) InsertData(sd *models.Log) error {
	res := d.DB.Save(sd)
	if res.Error != nil {
		log.Println("Error while creating entry in db", res.Error)
		return res.Error
	}
	return nil
}

func (d *UserStore) CreateUser(user *models.User) error {
	res := d.DB.Create(user)
	if res.Error != nil {
		log.Println("Error while creating entry in db", res.Error)
		return res.Error
	}
	return nil
}
