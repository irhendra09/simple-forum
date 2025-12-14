package model

import "time"

type Posts struct {
	ID           int64     `db:"id"`
	UserID       int64     `db:"user_id"`
	PostTitle    string    `db:"post_title"`
	PostContent  string    `db:"post_content"`
	PostHashtags string    `db:"post_hashtags"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	CreatedBy    string    `db:"created_by"`
	UpdatedBy    string    `db:"updated_by"`
}

type (
	CreatePostRequest struct {
		PostTitle    string   `json:"postTitle"`
		PostContent  string   `json:"postContent"`
		PostHashtags []string `json:"postHashtags"`
	}
)

type (
	GetAllPostResponse struct {
		Data       []Post     `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	Post struct {
		ID           int64    `json:"id"`
		UserID       int64    `json:"userID"`
		Username     string   `json:"username"`
		PostTitle    string   `json:"postTitle"`
		PostContent  string   `json:"postContent"`
		PostHashtags []string `json:"postHashtags"`
		IsLiked      bool     `json:"isLiked"`
	}

	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}

	GetPostResponse struct {
		PostDetail Post      `json:"postDetail"`
		LikeCount  int64     `json:"likeCount"`
		Comments   []Comment `json:"comments"`
	}

	Comment struct {
		ID             int64  `json:"id"`
		UserID         int64  `json:"user_id"`
		Username       string `json:"username"`
		CommentContent string `json:"commentContent"`
	}
)

type UserActivity struct {
	ID        int64     `db:"id"`
	PostID    int64     `db:"post_id"`
	UserID    int64     `db:"user_id"`
	IsLiked   bool      `db:"is_liked"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedBy string    `db:"created_by"`
	UpdatedBy string    `db:"updated_by"`
}

type (
	UserActivityRequest struct {
		IsLiked bool `json:"isLiked"`
	}
)

type PaginatedResult[T any] struct {
	Data       []T   `json:"data"`
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalPages int   `json:"total_pages"`
}
