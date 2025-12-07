package repository

import (
	"donedev.com/simple-forum/internal/configs"
	"donedev.com/simple-forum/internal/model"
)

func CreateComment(comment *model.Comments) error {
	return configs.ConnectDB().Create(comment).Error
}
