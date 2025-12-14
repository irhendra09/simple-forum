package repository

import (
	"donedev.com/simple-forum/internal/interfaces"
	"donedev.com/simple-forum/internal/model"
	"gorm.io/gorm"
)

// package-level helpers removed to enforce using repository instances

type GormCommentRepository struct {
	db *gorm.DB
}

func NewGormCommentRepository(db *gorm.DB) *GormCommentRepository {
	return &GormCommentRepository{db: db}
}

func (r *GormCommentRepository) CreateComment(comment *model.Comments) error {
	return r.db.Create(comment).Error
}

func (r *GormCommentRepository) GetCommentByPostId(postId int64) ([]model.Comment, error) {
	var result []model.Comment

	err := r.db.
		Table("comments c").
		Select("c.id, c.user_id, c.comment_content, u.username").
		Joins("JOIN users u ON c.user_id = u.id").
		Where("c.post_id = ?", postId).
		Scan(&result).Error
	if err != nil {
		return nil, err
	}
	comments := make([]model.Comment, 0, len(result))
	for _, r := range result {
		comments = append(comments, model.Comment{
			ID:             r.ID,
			UserID:         r.UserID,
			CommentContent: r.CommentContent,
			Username:       r.Username,
		})
	}
	return comments, nil
}

// ensure compile-time that GormCommentRepository implements interfaces.CommentRepository
var _ interfaces.CommentRepository = (*GormCommentRepository)(nil)
