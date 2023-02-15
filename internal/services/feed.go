package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	kafkaConfig "social-network/config/kafka"
	"social-network/internal/entities"
	"social-network/internal/handlers"
	"social-network/internal/repositories"
	"social-network/up"
	"social-network/utils/golibs/idutil"
	"social-network/utils/kafka"
	"social-network/utils/xerror"
)

var _ up.FeedService = &FeedService{}

type FeedService struct {
	feedRepo      *repositories.FeedRepository
	kafkaProducer *kafka.KafkaProducer
}

func NewFeedService(db *sql.DB, kafkaProducer *kafka.KafkaProducer) *FeedService {
	return &FeedService{
		feedRepo:      repositories.NewFeedRepository(db),
		kafkaProducer: kafkaProducer,
	}
}

func (s *FeedService) Create(ctx context.Context, req *up.CreateFeedRequest) (*up.CreateFeedResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	currentAccount, _ := accountIDFromCtx(ctx)

	now := time.Now()
	feed := &entities.Feed{
		FeedID:    idutil.NewID(),
		AccountID: currentAccount,
		Message:   req.Message,
		ImageUrl:  req.ImageUrl,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	if err := s.feedRepo.Create(ctx, feed); err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.feedRepository.Create: %w", err))
	}

	// send message to kafka
	body, err := handlers.NewFeedCreated(feed)
	if err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("handlers.NewFeedCreated: %w", err))
	}

	if err := s.kafkaProducer.SendMessage(ctx, string(kafkaConfig.FeedTopic), body); err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.kafkaProducer.SendMessage: %w", err))
	}

	return &up.CreateFeedResponse{
		FeedID: feed.FeedID,
	}, nil
}

func (s *FeedService) Update(ctx context.Context, req *up.UpdateFeedRequest) (*up.UpdateFeedResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	currentAccount, _ := accountIDFromCtx(ctx)

	now := time.Now()
	feed := &entities.Feed{
		FeedID:    req.FeedID,
		AccountID: currentAccount,
		Message:   req.Message,
		ImageUrl:  req.ImageUrl,
		UpdatedAt: &now,
	}

	if err := s.feedRepo.Update(ctx, feed); err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.feedRepository.Update: %w", err))
	}

	return &up.UpdateFeedResponse{}, nil
}

func (s *FeedService) Get(ctx context.Context, req *up.GetFeedRequest) (*up.GetFeedResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	currentAccount, _ := accountIDFromCtx(ctx)

	feed, err := s.feedRepo.FindByFeedIDAndAccountID(ctx, req.FeedID, currentAccount)
	if err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.feedRepository.FindByFeedIDAndAccountID: %w", err))
	}

	return &up.GetFeedResponse{
		Feed: s.convertFeedEntToFeedUp(feed),
	}, nil
}

func (s *FeedService) List(ctx context.Context, req *up.ListFeedsRequest) (*up.ListFeedsResponse, error) {
	listFeedsArgs := &repositories.ListFeedsArgs{}
	if req.OnlyCurrentAccount {
		currentAccount, _ := accountIDFromCtx(ctx)

		listFeedsArgs.AccountID = &currentAccount
	}

	feeds, err := s.feedRepo.List(ctx, listFeedsArgs)
	if err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.feedRepo.List: %w", err))
	}

	feedsUp := make([]*up.Feed, 0, len(feeds))
	for _, feed := range feeds {
		feedsUp = append(feedsUp, s.convertFeedEntToFeedUp(feed))
	}

	return &up.ListFeedsResponse{
		Feeds: feedsUp,
	}, nil
}

func (s *FeedService) Delete(ctx context.Context, req *up.DeleteFeedRequest) (*up.DeleteFeedResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	currentAccount, _ := accountIDFromCtx(ctx)

	err := s.feedRepo.Delete(ctx, req.FeedID, currentAccount)
	if err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.feedRepository.Delete: %w", err))
	}

	return &up.DeleteFeedResponse{}, nil
}

func (s *FeedService) convertFeedEntToFeedUp(feedEnt *entities.Feed) *up.Feed {
	feedUp := &up.Feed{
		FeedID:    feedEnt.FeedID,
		AccountID: feedEnt.AccountID,
		Message:   feedEnt.Message,
		ImageUrl:  feedEnt.ImageUrl,
	}
	if feedEnt.CreatedAt != nil {
		feedUp.CreatedAt = *feedEnt.CreatedAt
	}
	if feedEnt.UpdatedAt != nil {
		feedUp.UpdatedAt = *feedEnt.UpdatedAt
	}

	return feedUp
}
