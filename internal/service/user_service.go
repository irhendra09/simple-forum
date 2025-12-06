package service

import (
	"errors"

	"donedev.com/simple-forum/internal/model"
	"donedev.com/simple-forum/internal/repository"
	"donedev.com/simple-forum/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(user *model.SignUpRequest) (*model.Users, error) {
	if repository.IsEmailTaken(user.Email) {
		return nil, errors.New("Email already taken")
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

	err = repository.CreateUser(&users)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func Login(request *model.LoginRequest) (*model.LoginResponse, error) {
	user, err := repository.GetUsersByEmail(request.Email)
	if err != nil {
		return nil, errors.New("email atau password salah")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return nil, errors.New("email atau password salah")
	}

	accessToken, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
