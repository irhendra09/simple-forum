package repository

import (
	"errors"

	apperrors "donedev.com/simple-forum/internal/errors"
	"donedev.com/simple-forum/internal/interfaces"
	"donedev.com/simple-forum/internal/model"
	"gorm.io/gorm"
)

// package-level helpers removed to enforce using repository instances

// GormUserRepository implements interfaces.UserRepository using gorm
type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) CreateUser(user *model.Users) error {
	return r.db.Create(user).Error
}

func (r *GormUserRepository) GetUsersByEmail(email string) (*model.Users, error) {
	var users model.Users
	err := r.db.Where("email = ?", email).First(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrUnauthorized
		}
		return nil, err
	}
	return &users, nil
}

func (r *GormUserRepository) IsEmailTaken(email string) bool {
	var count int64
	r.db.Model(&model.Users{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// ensure compile-time that GormUserRepository implements interfaces.UserRepository
var _ interfaces.UserRepository = (*GormUserRepository)(nil)
