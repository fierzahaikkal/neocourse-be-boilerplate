package repository

import (
	"github.com/fierzahaikkal/neocourse-be-boilerplate-golang/internal/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB  *gorm.DB
	log *log.Logger
}

func NewUserRepository(db *gorm.DB, log *log.Logger) *UserRepository {
	return &UserRepository{
		DB:  db,
		log: log,
	}
}

func (r *UserRepository) Register(user *entity.User) error {
	var existingUser entity.User

	err := r.DB.Where("email = ?", user.Email).First(&existingUser).Error
	if err != nil {
		return err
	}

	err = r.DB.Where("username = ?", user.Username).First(&existingUser).Error
	if err != nil {
		return err
	}

	err = r.DB.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}
