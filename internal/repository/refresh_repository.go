package repository

import (
	"time"

	"donedev.com/simple-forum/internal/interfaces"
	"donedev.com/simple-forum/internal/model"
	"gorm.io/gorm"
)

type GormRefreshRepository struct {
	db *gorm.DB
}

func NewGormRefreshRepository(db *gorm.DB) *GormRefreshRepository {
	return &GormRefreshRepository{db: db}
}

func (r *GormRefreshRepository) CreateRefreshToken(token *model.RefreshToken) error {
	if token.CreatedAt.IsZero() {
		token.CreatedAt = time.Now()
	}
	return r.db.Create(token).Error
}

func (r *GormRefreshRepository) GetRefreshTokenByToken(token string) (*model.RefreshToken, error) {
	var t model.RefreshToken
	err := r.db.Where("token = ?", token).First(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *GormRefreshRepository) RevokeRefreshTokenByToken(token string) error {
	return r.db.Model(&model.RefreshToken{}).Where("token = ?", token).Update("revoked", true).Error
}

// compile-time assertion
var _ interfaces.TokenRepository = (*GormRefreshRepository)(nil)
