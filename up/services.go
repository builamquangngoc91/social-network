package up

import (
	"context"
)

type AccountService interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Follow(context.Context, *FollowRequest) (*FollowResponse, error)
	UnFollow(context.Context, *UnFollowRequest) (*UnFollowResponse, error)
}

type FeedService interface {
	Create(context.Context, *CreateFeedRequest) (*CreateFeedResponse, error)
	Update(context.Context, *UpdateFeedRequest) (*UpdateFeedResponse, error)
	Get(context.Context, *GetFeedRequest) (*GetFeedResponse, error)
	List(context.Context, *ListFeedsRequest) (*ListFeedsResponse, error)
	Search(context.Context, *SearchFeedsRequest) (*SearchFeedsResponse, error)
	Delete(context.Context, *DeleteFeedRequest) (*DeleteFeedResponse, error)
}

type CommentService interface {
	Create(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	Update(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error)
	List(context.Context, *ListCommentsRequest) (*ListCommentsResponse, error)
	Delete(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error)
}

// type MovieService interface {
// 	Create(context.Context, *CreateMovieRequest) (*CreateMovieResponse, error)
// 	GetMovieByUser(context.Context, *GetMovieByUserRequest) (*GetMovieByUserResponse, error)
// 	ListMoviesByUser(context.Context, *ListMoviesByUserRequest) (*ListMoviesByUserResponse, error)
// 	ListMovies(context.Context, *ListMoviesRequest) (*ListMoviesResponse, error)
// }
