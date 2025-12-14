package interfaces

import "donedev.com/simple-forum/internal/model"

type UserService interface {
	Register(user *model.SignUpRequest) (*model.Users, error)
	Login(request *model.LoginRequest) (*model.LoginResponse, error)
	RefreshToken(refreshToken string) (*model.LoginResponse, error)
	Logout(refreshToken string) error
}

type PostService interface {
	CreatePost(req *model.CreatePostRequest, userId int64) error
	GetAllPosts(page, pageSize int) (*model.PaginatedResult[model.Posts], error)
	GetPostById(id int64) (*model.GetPostResponse, error)
	UpsertUserActivity(request model.UserActivityRequest, postId, userId int64) error
}

type CommentService interface {
	CreateComment(postID, userId int64, comments *model.CreateCommentRequest) error
}

type TokenService interface {
	GenerateToken(userID int64) (string, error)
	GenerateRefreshToken(userID int64) (string, error)
	ParseToken(token string) (int64, error)
}
