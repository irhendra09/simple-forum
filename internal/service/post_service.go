package service

import (
	"log"
	"strconv"
	"strings"
	"time"

	"donedev.com/simple-forum/internal/model"
	"donedev.com/simple-forum/internal/repository"
)

func CreatePost(req *model.CreatePostRequest, userId int64) error {
	postHastag := strings.Join(req.PostHashtags, ",")
	post := model.Posts{
		PostHashtags: postHastag,
		PostTitle:    req.PostTitle,
		PostContent:  req.PostContent,
		UserID:       userId,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    strconv.FormatInt(userId, 10),
		UpdatedBy:    strconv.FormatInt(userId, 10),
	}
	err := repository.CreatePost(&post)

	if err != nil {
		log.Println("error creating post:", err)
		return err
	}
	return nil
}
