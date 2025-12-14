package service

import (
	"strconv"
	"time"

	"donedev.com/simple-forum/internal/interfaces"
	"donedev.com/simple-forum/internal/model"
	"github.com/rs/zerolog/log"
)

var CommentService interfaces.CommentService

type commentService struct {
	repo interfaces.CommentRepository
}

func NewCommentService(repo interfaces.CommentRepository) interfaces.CommentService {
	s := &commentService{repo: repo}
	CommentService = s
	return s
}

func (s *commentService) CreateComment(postID, userId int64, comments *model.CreateCommentRequest) error {
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
	err := s.repo.CreateComment(comment)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create comment")
		return err
	}
	return nil
}
