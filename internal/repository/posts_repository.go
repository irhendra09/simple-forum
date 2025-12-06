package repository

import (
	"donedev.com/simple-forum/internal/configs"
	"donedev.com/simple-forum/internal/model"
)

func CreatePost(posts *model.Posts) error {
	return configs.ConnectDB().Create(posts).Error
}
