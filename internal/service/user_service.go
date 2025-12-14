package service

import (
	"time"

	apperrors "donedev.com/simple-forum/internal/errors"
	"donedev.com/simple-forum/internal/interfaces"
	"donedev.com/simple-forum/internal/model"
	"golang.org/x/crypto/bcrypt"
)

var UserService interfaces.UserService

type userService struct {
	repo      interfaces.UserRepository
	token     interfaces.TokenService
	tokenRepo interfaces.TokenRepository
}

func NewUserService(repo interfaces.UserRepository, tokenSvc interfaces.TokenService, tokenRepo interfaces.TokenRepository) interfaces.UserService {
	s := &userService{repo: repo, token: tokenSvc, tokenRepo: tokenRepo}
	UserService = s
	return s
}

func (s *userService) Register(user *model.SignUpRequest) (*model.Users, error) {
	if s.repo.IsEmailTaken(user.Email) {
		return nil, apperrors.ErrConflict
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	users := model.Users{
		Email:    user.Email,
		Username: user.Username,
		Password: string(hashed),
	}

	err = s.repo.CreateUser(&users)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (s *userService) Login(request *model.LoginRequest) (*model.LoginResponse, error) {
	user, err := s.repo.GetUsersByEmail(request.Email)
	if err != nil {
		return nil, apperrors.ErrUnauthorized
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return nil, apperrors.ErrUnauthorized
	}

	accessToken, err := s.token.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.token.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	// persist refresh token
	expiresAt := time.Now().Add(24 * time.Hour)
	rt := &model.RefreshToken{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: expiresAt,
		Revoked:   false,
	}
	if err := s.tokenRepo.CreateRefreshToken(rt); err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *userService) RefreshToken(refreshToken string) (*model.LoginResponse, error) {
	// validate token structure and get user id
	userID, err := s.token.ParseToken(refreshToken)
	if err != nil {
		return nil, err
	}

	// check persisted token
	stored, err := s.tokenRepo.GetRefreshTokenByToken(refreshToken)
	if err != nil {
		return nil, err
	}
	if stored.Revoked {
		return nil, apperrors.ErrUnauthorized
	}
	if stored.ExpiresAt.Before(time.Now()) {
		return nil, apperrors.ErrUnauthorized
	}

	// revoke old token
	if err := s.tokenRepo.RevokeRefreshTokenByToken(refreshToken); err != nil {
		return nil, err
	}

	// generate new tokens and persist new refresh token
	accessToken, err := s.token.GenerateToken(userID)
	if err != nil {
		return nil, err
	}
	newRefreshToken, err := s.token.GenerateRefreshToken(userID)
	if err != nil {
		return nil, err
	}
	expiresAt := time.Now().Add(24 * time.Hour)
	newRT := &model.RefreshToken{
		Token:     newRefreshToken,
		UserID:    userID,
		ExpiresAt: expiresAt,
		Revoked:   false,
	}
	if err := s.tokenRepo.CreateRefreshToken(newRT); err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *userService) Logout(refreshToken string) error {
	// revoke the provided refresh token
	// parse token to ensure it's valid structure (optional)
	if _, err := s.token.ParseToken(refreshToken); err != nil {
		return err
	}
	return s.tokenRepo.RevokeRefreshTokenByToken(refreshToken)
}
