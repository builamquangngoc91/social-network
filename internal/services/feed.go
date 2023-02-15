package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	kafkaConfig "social-network/config/kafka"
	"social-network/internal/entities"
	"social-network/internal/handlers"
	"social-network/internal/repositories"
	"social-network/up"
	"social-network/utils/elasticsearch"
	"social-network/utils/golibs/idutil"
	"social-network/utils/kafka"
	"social-network/utils/xerror"

	"github.com/olivere/elastic/v7"
	"github.com/redis/go-redis/v9"
)

var _ up.FeedService = &FeedService{}

type FeedService struct {
	feedRepo      *repositories.FeedRepository
	esClient      *elasticsearch.ElasticClient
	kafkaProducer *kafka.KafkaProducer
	rd            *redis.Client
}

func NewFeedService(db *sql.DB, kafkaProducer *kafka.KafkaProducer, esClient *elasticsearch.ElasticClient, rd *redis.Client) *FeedService {
	return &FeedService{
		feedRepo:      repositories.NewFeedRepository(db),
		esClient:      esClient,
		kafkaProducer: kafkaProducer,
		rd:            rd,
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
		Tag:       req.Tag,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	if err := s.feedRepo.Create(ctx, feed); err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.feedRepository.Create: %w", err))
	}

	// send message to kafka
	body, err := handlers.MarshalFeed(feed)
	if err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("handlers.MarshalFeed: %w", err))
	}

	if err := s.kafkaProducer.SendMessage(ctx, string(kafkaConfig.FeedTopic), body); err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.kafkaProducer.SendMessage: %w", err))
	}

	if err := s.rd.ZIncrBy(ctx, "leaderboard", 1, currentAccount).Err(); err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("cannot update leaderboard"))
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

func (s *FeedService) Search(ctx context.Context, req *up.SearchFeedsRequest) (*up.SearchFeedsResponse, error) {
	switch req.Field {
	case "tag":
		searchParams := &elasticsearch.SearchParams{
			Query: elastic.NewMatchQuery(req.Field, req.Value),
			SortBy: []elastic.Sorter{
				elastic.NewFieldSort("message").Desc().Sorter,
			},
			Index: "feed",
		}
		if req.OffsetPaging != nil {
			searchParams.Limit = req.OffsetPaging.Limit
			searchParams.Offset = req.OffsetPaging.Offset
		}

		var (
			feed entities.Feed
		)
		result, err := s.esClient.Search(ctx, searchParams)
		if err != nil {
			return nil, fmt.Errorf("s.esClient.Search: %w", err)
		}

		feedsUp := make([]*up.Feed, 0, len(result.Hits.Hits))
		for _, hit := range result.Hits.Hits {
			err := json.Unmarshal(hit.Source, &feed)
			if err != nil {
				fmt.Println("[Getting Feeds][Unmarshal] Err=", err)
			}

			feedsUp = append(feedsUp, s.convertFeedEntToFeedUp(&feed))
		}

		return &up.SearchFeedsResponse{
			Feeds: feedsUp,
		}, nil
	case "message":
		feeds, err := s.feedRepo.Search(ctx, req.Value)
		if err != nil {
			return nil, fmt.Errorf("s.feedRepo.Search: %w", err)
		}

		feedsUp := make([]*up.Feed, 0, len(feeds))
		for _, feed := range feeds {
			feedsUp = append(feedsUp, s.convertFeedEntToFeedUp(feed))
		}

		return &up.SearchFeedsResponse{
			Feeds: feedsUp,
		}, nil
	}

	return &up.SearchFeedsResponse{}, nil
}

func (s *FeedService) GetLeaderBoard(ctx context.Context, req *up.GetLeaderBoardRequest) (*up.GetLeaderBoardResponse, error) {
	top := req.Top
	if top == 0 {
		top = 10
	}

	zs, err := s.rd.ZRangeArgsWithScores(ctx, redis.ZRangeArgs{
		Rev:   true,
		Start: 0,
		Stop:  top - 1,
		Key:   "leaderboard",
	}).Result()
	if err != nil {
		return nil, fmt.Errorf("s.rd.ZRangeArgsWithScores: %w", err)
	}

	rows := make([]*up.Row, 0, len(zs))
	for _, z := range zs {
		rows = append(rows, &up.Row{
			AccountID:     z.Member.(string),
			NumberOfFeeds: z.Score,
		})
	}
	return &up.GetLeaderBoardResponse{
		Rows: rows,
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
		Tag:       feedEnt.Tag,
	}
	if feedEnt.CreatedAt != nil {
		feedUp.CreatedAt = *feedEnt.CreatedAt
	}
	if feedEnt.UpdatedAt != nil {
		feedUp.UpdatedAt = *feedEnt.UpdatedAt
	}

	return feedUp
}
