package up

import (
	"context"
)

type AccountService interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Follow(context.Context, *FollowAccountRequest) (*FollowAccountResponse, error)
	UnFollow(context.Context, *UnFollowAccountRequest) (*UnFollowAccountResponse, error)
}

type FeedService interface {
	Create(context.Context, *CreateFeedRequest) (*CreateFeedResponse, error)
	Update(context.Context, *UpdateFeedRequest) (*UpdateFeedRequest, error)
	Get(context.Context, *GetFeedRequest) (*GetFeedResponse, error)
}

// type MovieService interface {
// 	Create(context.Context, *CreateMovieRequest) (*CreateMovieResponse, error)
// 	GetMovieByUser(context.Context, *GetMovieByUserRequest) (*GetMovieByUserResponse, error)
// 	ListMoviesByUser(context.Context, *ListMoviesByUserRequest) (*ListMoviesByUserResponse, error)
// 	ListMovies(context.Context, *ListMoviesRequest) (*ListMoviesResponse, error)
// }
