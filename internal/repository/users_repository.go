package repository

import (
	"errors"

	"donedev.com/simple-forum/internal/configs"
	"donedev.com/simple-forum/internal/model"
	"gorm.io/gorm"
)

func CreateUser(user *model.Users) error {
	return configs.ConnectDB().Create(user).Error
}

func GetUsersByEmail(email string) (*model.Users, error) {
	var users model.Users
	err := configs.ConnectDB().Where("email = ?", email).First(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("email atau password salah")
		}
		return nil, err
	}
	return &users, nil
}

func IsEmailTaken(email string) bool {
	var count int64
	configs.ConnectDB().Model(&model.Users{}).Where("email = ?", email).Count(&count)
	return count > 0
}
