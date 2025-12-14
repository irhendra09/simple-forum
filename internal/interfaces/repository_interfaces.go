package interfaces

import "donedev.com/simple-forum/internal/model"

type UserRepository interface {
	CreateUser(user *model.Users) error
	GetUsersByEmail(email string) (*model.Users, error)
	IsEmailTaken(email string) bool
}

type PostRepository interface {
	CreatePost(posts *model.Posts) error
	GetAllPosts(page, pageSize int) (*model.PaginatedResult[model.Posts], error)
	GetPostById(postId int64) (*model.Post, error)
	GetCommentByPostId(postId int64) ([]model.Comment, error)
	GetUserActivity(user model.UserActivity) (*model.UserActivity, error)
	CreateUserActivity(userActivity *model.UserActivity) error
	UpdateUserActivity(userActivity *model.UserActivity) error
	CountLikeByPostId(id int64) (int64, error)
}

type CommentRepository interface {
	CreateComment(comment *model.Comments) error
	GetCommentByPostId(postId int64) ([]model.Comment, error)
}

type TokenRepository interface {
	CreateRefreshToken(token *model.RefreshToken) error
	GetRefreshTokenByToken(token string) (*model.RefreshToken, error)
	RevokeRefreshTokenByToken(token string) error
}
