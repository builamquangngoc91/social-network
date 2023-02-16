package server

import (
	"database/sql"
	"net/http"

	"social-network/internal/services"
	"social-network/utils/acl"
	"social-network/utils/elasticsearch"
	"social-network/utils/kafka"

	"github.com/redis/go-redis/v9"
)

type FeedServer struct {
	jwtKey string

	feedService    *services.FeedService
	commentService *services.CommentService

	kafkaProducer *kafka.KafkaProducer
	esClient      *elasticsearch.ElasticClient
	rd            *redis.Client

	acList map[string]map[string]acl.Decl

	db *sql.DB
}

func NewFeedServer(
	db *sql.DB, JWTKey string,
	kafkaProducer *kafka.KafkaProducer, esClient *elasticsearch.ElasticClient,
	rd *redis.Client,
) *FeedServer {
	feedService := services.NewFeedService(db, kafkaProducer, esClient, rd)
	commentService := services.NewCommentService(db, kafkaProducer)

	return &FeedServer{
		jwtKey:         JWTKey,
		kafkaProducer:  kafkaProducer,
		feedService:    feedService,
		commentService: commentService,
		esClient:       esClient,
		rd:             rd,
		db:             db,
		acList: map[string]map[string]acl.Decl{
			"/api/v1/createFeed": {
				http.MethodPost: acl.Decl{
					HandlerFunc:  feedService.Create,
					Auth:         acl.User,
					ResponseType: acl.JSON,
				},
			},
			"/api/v1/updateFeed": {
				http.MethodPost: acl.Decl{
					HandlerFunc:  feedService.Update,
					Auth:         acl.User,
					ResponseType: acl.JSON,
				},
			},
			"/api/v1/getFeed": {
				http.MethodPost: acl.Decl{
					HandlerFunc:  feedService.Get,
					Auth:         acl.User,
					ResponseType: acl.JSON,
				},
			},
			"/api/v1/listFeeds": {
				http.MethodPost: acl.Decl{
					HandlerFunc:  feedService.List,
					Auth:         acl.User,
					ResponseType: acl.JSON,
				},
			},
			"/api/v1/deleteFeed": {
				http.MethodPost: acl.Decl{
					HandlerFunc:  feedService.Delete,
					Auth:         acl.User,
					ResponseType: acl.JSON,
				},
			},
			"/api/v1/searchFeeds": {
				http.MethodPost: acl.Decl{
					HandlerFunc:  feedService.Search,
					Auth:         acl.User,
					ResponseType: acl.JSON,
				},
			},
			"/api/v1/createComment": {
				http.MethodPost: acl.Decl{
					HandlerFunc:  commentService.Create,
					Auth:         acl.User,
					ResponseType: acl.JSON,
				},
			},
			"/api/v1/updateComment": {
				http.MethodPost: acl.Decl{
					HandlerFunc:  commentService.Update,
					Auth:         acl.User,
					ResponseType: acl.JSON,
				},
			},
			"/api/v1/listComments": {
				http.MethodPost: acl.Decl{
					HandlerFunc:  commentService.List,
					Auth:         acl.User,
					ResponseType: acl.JSON,
				},
			},
			"/api/v1/deleteComment": {
				http.MethodPost: acl.Decl{
					HandlerFunc:  commentService.Delete,
					Auth:         acl.User,
					ResponseType: acl.JSON,
				},
			},
			"/api/v1/getFeedLeaderBoard": {
				http.MethodPost: acl.Decl{
					HandlerFunc:  feedService.GetLeaderBoard,
					Auth:         acl.None,
					ResponseType: acl.JSON,
				},
			},
		},
	}
}

func (s *FeedServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	acl.ServeHTTP(resp, req, s.acList, s.jwtKey)
}
