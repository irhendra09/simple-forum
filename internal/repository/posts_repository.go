package repository

import (
	"errors"
	"strings"

	apperrors "donedev.com/simple-forum/internal/errors"
	"donedev.com/simple-forum/internal/interfaces"
	"donedev.com/simple-forum/internal/model"
	"gorm.io/gorm"
)

// package-level helpers removed to enforce using repository instances

// GormPostRepository implements interfaces.PostRepository using gorm
type GormPostRepository struct {
	db *gorm.DB
}

func NewGormPostRepository(db *gorm.DB) *GormPostRepository {
	return &GormPostRepository{db: db}
}

func (r *GormPostRepository) CreatePost(posts *model.Posts) error {
	return r.db.Create(posts).Error
}

func (r *GormPostRepository) GetAllPosts(page, pageSize int) (*model.PaginatedResult[model.Posts], error) {
	var postData []model.Posts
	var count int64

	if err := r.db.Model(&model.Posts{}).Count(&count).Error; err != nil {
		return nil, err
	}
	offset := (page - 1) * pageSize
	if err := r.db.Limit(pageSize).Offset(offset).Order("created_at DESC").Find(&postData).Error; err != nil {
		return nil, err
	}
	countPages := int((count + int64(pageSize) - 1) / int64(pageSize))
	result := model.PaginatedResult[model.Posts]{
		Data:       postData,
		Page:       page,
		PageSize:   pageSize,
		Total:      count,
		TotalPages: countPages,
	}
	return &result, nil
}

func (r *GormPostRepository) GetPostById(postId int64) (*model.Post, error) {
	// scan into a temporary struct where PostHashtags is a string
	var tmp struct {
		ID           int64  `gorm:"column:id"`
		UserID       int64  `gorm:"column:user_id"`
		Username     string `gorm:"column:username"`
		PostTitle    string `gorm:"column:post_title"`
		PostContent  string `gorm:"column:post_content"`
		PostHashtags string `gorm:"column:post_hashtags"`
		IsLiked      bool   `gorm:"column:is_liked"`
	}

	err := r.db.
		Table("posts p").
		Select("p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hashtags, uv.is_liked").
		Joins("JOIN users u ON p.user_id = u.id").
		Joins("JOIN user_activities uv ON uv.post_id = p.id").
		Where("p.id = ?", postId).
		Scan(&tmp).Error
	if err != nil {
		return nil, err
	}

	// convert comma-separated hashtags into []string
	var hashtags []string
	if tmp.PostHashtags != "" {
		// split and trim spaces
		parts := strings.Split(tmp.PostHashtags, ",")
		for _, p := range parts {
			s := strings.TrimSpace(p)
			if s != "" {
				hashtags = append(hashtags, s)
			}
		}
	}

	result := model.Post{
		ID:           tmp.ID,
		UserID:       tmp.UserID,
		Username:     tmp.Username,
		PostTitle:    tmp.PostTitle,
		PostContent:  tmp.PostContent,
		PostHashtags: hashtags,
		IsLiked:      tmp.IsLiked,
	}
	return &result, nil
}

func (r *GormPostRepository) GetUserActivity(user model.UserActivity) (*model.UserActivity, error) {
	var userActivity model.UserActivity
	err := r.db.Where("user_id = ? AND post_id = ?", user.UserID, user.PostID).First(&userActivity).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &userActivity, err
}

func (r *GormPostRepository) CreateUserActivity(userActivity *model.UserActivity) error {
	return r.db.Create(userActivity).Error
}

func (r *GormPostRepository) UpdateUserActivity(userActivity *model.UserActivity) error {
	return r.db.Save(userActivity).Error
}

func (r *GormPostRepository) CountLikeByPostId(id int64) (int64, error) {
	var count int64
	err := r.db.Model(&model.UserActivity{}).Where("post_id = ? AND is_liked = ?", id, true).Count(&count).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, apperrors.ErrNotFound
		}
		return 0, err
	}
	return count, nil
}

func (r *GormPostRepository) GetCommentByPostId(postId int64) ([]model.Comment, error) {
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

// ensure compile-time that GormPostRepository implements interfaces.PostRepository
var _ interfaces.PostRepository = (*GormPostRepository)(nil)
