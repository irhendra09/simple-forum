package service

import (
	"strconv"
	"strings"
	"time"

	apperrors "donedev.com/simple-forum/internal/errors"
	"donedev.com/simple-forum/internal/interfaces"
	"donedev.com/simple-forum/internal/model"
	"github.com/rs/zerolog/log"
)

var PostService interfaces.PostService

type postService struct {
	repo interfaces.PostRepository
}

func NewPostService(repo interfaces.PostRepository) interfaces.PostService {
	s := &postService{repo: repo}
	PostService = s
	return s
}

func (s *postService) CreatePost(req *model.CreatePostRequest, userId int64) error {
	postHashtag := strings.Join(req.PostHashtags, ",")
	post := model.Posts{
		PostHashtags: postHashtag,
		PostTitle:    req.PostTitle,
		PostContent:  req.PostContent,
		UserID:       userId,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    strconv.FormatInt(userId, 10),
		UpdatedBy:    strconv.FormatInt(userId, 10),
	}
	err := s.repo.CreatePost(&post)

	if err != nil {
		log.Error().Err(err).Msg("error creating post:")
		return err
	}
	return nil
}

func (s *postService) GetAllPosts(page, pageSize int) (*model.PaginatedResult[model.Posts], error) {
	if pageSize <= 0 {
		pageSize = 10
	}
	if page < 1 {
		page = 1
	}

	return s.repo.GetAllPosts(page, pageSize)
}

func (s *postService) GetPostById(id int64) (*model.GetPostResponse, error) {
	postDetail, err := s.repo.GetPostById(id)
	if err != nil {
		log.Error().Err(err).Msg("error getting post by id:")
		return nil, err
	}
	likeCount, err := s.repo.CountLikeByPostId(id)
	if err != nil {
		log.Error().Err(err).Msg("error count like to database")
		return nil, err
	}

	comments, err := s.repo.GetCommentByPostId(id)
	if err != nil {
		log.Error().Err(err).Msg("error get comment by post id:")
		return nil, err
	}

	return &model.GetPostResponse{
		PostDetail: model.Post{
			ID:           postDetail.ID,
			UserID:       postDetail.UserID,
			Username:     postDetail.Username,
			PostTitle:    postDetail.PostTitle,
			PostContent:  postDetail.PostContent,
			PostHashtags: postDetail.PostHashtags,
			IsLiked:      postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comments:  comments,
	}, nil
}

func (s *postService) UpsertUserActivity(request model.UserActivityRequest, postId, userId int64) error {
	now := time.Now()
	activity := model.UserActivity{
		PostID:    postId,
		UserID:    userId,
		IsLiked:   request.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userId, 10),
		UpdatedBy: strconv.FormatInt(userId, 10),
	}
	userActivity, err := s.repo.GetUserActivity(activity)
	if err != nil {
		log.Error().Err(err).Msg("Error getting user activity")
		return err
	}
	if userActivity == nil {
		if !request.IsLiked {
			return apperrors.ErrBadRequest
		}
		err = s.repo.CreateUserActivity(&activity)
	} else {
		err = s.repo.UpdateUserActivity(userActivity)
	}
	if err != nil {
		log.Error().Err(err).Msg("Error creating and updating user activity")
		return err
	}
	return nil
}
