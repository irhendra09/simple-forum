package repository

import (
	"donedev.com/simple-forum/internal/configs"
	"donedev.com/simple-forum/internal/model"
)

func CreatePost(posts *model.Posts) error {
	return configs.ConnectDB().Create(posts).Error
}

func GetPostById(postId int64) error {
	var post model.Posts
	return configs.ConnectDB().Where("id = ?", postId).First(&post).Error
}
