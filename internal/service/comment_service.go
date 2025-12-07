package service

import (
	"strconv"
	"time"

	"donedev.com/simple-forum/internal/model"
	"donedev.com/simple-forum/internal/repository"
	"github.com/rs/zerolog/log"
)

func CreateComment(postID, userId int64, comments *model.CreateCommentRequest) error {
	now := time.Now()

	comment := &model.Comments{
		PostID:         postID,
		UserID:         userId,
		CommentContent: comments.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      strconv.FormatInt(userId, 10),
		UpdatedBy:      strconv.FormatInt(userId, 10),
	}
	err := repository.CreateComment(comment)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create comment")
		return err
	}
	return nil
}
