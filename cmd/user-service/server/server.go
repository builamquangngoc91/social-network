package server

import (
	"database/sql"
	"net/http"

	"social-network/internal/services"
	"social-network/utils/acl"
	"social-network/utils/elasticsearch"
	"social-network/utils/kafka"
)

type UserServer struct {
	jwtKey string

	accountService *services.AccountService

	kafkaProducer *kafka.KafkaProducer
	esClient      *elasticsearch.ElasticClient

	acList map[string]map[string]acl.Decl

	db *sql.DB
}

func NewUserServer(
	db *sql.DB, JWTKey string,
	kafkaProducer *kafka.KafkaProducer, esClient *elasticsearch.ElasticClient,
) *UserServer {
	accountService := services.NewAccountService(db, JWTKey, kafkaProducer)

	return &UserServer{
		jwtKey:         JWTKey,
		kafkaProducer:  kafkaProducer,
		accountService: accountService,
		esClient:       esClient,
		db:             db,
		acList: map[string]map[string]acl.Decl{
			"/api/v1/register": {
				http.MethodPost: acl.Decl{
					HandlerFunc:  accountService.Register,
					Auth:         acl.None,
					ResponseType: acl.JSON,
				},
			},
			"/api/v1/login": {
				http.MethodPost: acl.Decl{
					HandlerFunc:  accountService.Login,
					Auth:         acl.None,
					ResponseType: acl.JSON,
				},
			},
			"/api/v1/followAccount": {
				http.MethodPost: acl.Decl{
					HandlerFunc:  accountService.Follow,
					Auth:         acl.User,
					ResponseType: acl.JSON,
				},
			},
			"/api/v1/unfollowAccount": {
				http.MethodPost: acl.Decl{
					HandlerFunc:  accountService.UnFollow,
					Auth:         acl.User,
					ResponseType: acl.JSON,
				},
			},
		},
	}
}

func (s *UserServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	acl.ServeHTTP(resp, req, s.acList, s.jwtKey)
}
