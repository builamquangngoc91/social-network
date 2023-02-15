package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"social-network/utils/kafka"
	"social-network/utils/xerror"

	"github.com/dgrijalva/jwt-go"
)

type (
	AuthType     int
	ResponseType int
)

const (
	None = AuthType(0)
	User = AuthType(1)

	JSON = ResponseType(0)
	HTML = ResponseType(1)
	File = ResponseType(2)
)

type Decl struct {
	Auth         AuthType
	HandlerFunc  interface{}
	ResponseType ResponseType
}

type RemiService struct {
	jwtKey         string
	kafkaProducer  *kafka.KafkaProducer
	accountService *AccountService
	feedService    *FeedService
	commentService *CommentService
	acl            map[string]map[string]Decl
}

func NewServices(db *sql.DB, JWTKey string, kafkaProducer *kafka.KafkaProducer) *RemiService {
	accountService := NewAccountService(db, JWTKey, kafkaProducer)
	feedService := NewFeedService(db, kafkaProducer)
	commentService := NewCommentService(db)

	return &RemiService{
		jwtKey:         JWTKey,
		kafkaProducer:  kafkaProducer,
		accountService: accountService,
		feedService:    feedService,
		commentService: commentService,
		acl: map[string]map[string]Decl{
			"/api/v1/register": {
				http.MethodPost: Decl{
					HandlerFunc:  accountService.Register,
					Auth:         None,
					ResponseType: JSON,
				},
			},
			"/api/v1/login": {
				http.MethodPost: Decl{
					HandlerFunc:  accountService.Login,
					Auth:         None,
					ResponseType: JSON,
				},
			},
			"/api/v1/followAccount": {
				http.MethodPost: Decl{
					HandlerFunc:  accountService.Follow,
					Auth:         User,
					ResponseType: JSON,
				},
			},
			"/api/v1/unfollowAccount": {
				http.MethodPost: Decl{
					HandlerFunc:  accountService.UnFollow,
					Auth:         User,
					ResponseType: JSON,
				},
			},
			"/api/v1/createFeed": {
				http.MethodPost: Decl{
					HandlerFunc:  feedService.Create,
					Auth:         User,
					ResponseType: JSON,
				},
			},
			"/api/v1/updateFeed": {
				http.MethodPost: Decl{
					HandlerFunc:  feedService.Update,
					Auth:         User,
					ResponseType: JSON,
				},
			},
			"/api/v1/getFeed": {
				http.MethodPost: Decl{
					HandlerFunc:  feedService.Get,
					Auth:         User,
					ResponseType: JSON,
				},
			},
			"/api/v1/listFeeds": {
				http.MethodPost: Decl{
					HandlerFunc:  feedService.List,
					Auth:         User,
					ResponseType: JSON,
				},
			},
			"/api/v1/deleteFeed": {
				http.MethodPost: Decl{
					HandlerFunc:  feedService.Delete,
					Auth:         User,
					ResponseType: JSON,
				},
			},
			"/api/v1/createComment": {
				http.MethodPost: Decl{
					HandlerFunc:  commentService.Create,
					Auth:         User,
					ResponseType: JSON,
				},
			},
			"/api/v1/updateComment": {
				http.MethodPost: Decl{
					HandlerFunc:  commentService.Update,
					Auth:         User,
					ResponseType: JSON,
				},
			},
			"/api/v1/listComments": {
				http.MethodPost: Decl{
					HandlerFunc:  commentService.List,
					Auth:         User,
					ResponseType: JSON,
				},
			},
			"/api/v1/deleteComment": {
				http.MethodPost: Decl{
					HandlerFunc:  commentService.Delete,
					Auth:         User,
					ResponseType: JSON,
				},
			},
		},
	}
}

func (s *RemiService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	handler, ok := s.acl[req.URL.Path]
	if !ok {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	decl, ok := handler[req.Method]
	if !ok {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	handlerFunc := decl.HandlerFunc
	auth := decl.Auth

	// authorization
	switch auth {
	case User:
		var ok bool
		req, ok = s.validToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}
	case None:
		// no-op
	}

	switch decl.ResponseType {
	case JSON:
		// Call functions of services
		// Todo: check function (type of arguments, outputs)
		typ := reflect.TypeOf(handlerFunc)
		funcArgs := reflect.New(typ.In(1).Elem()).Interface()

		// parse body and request values
		switch req.Method {
		case http.MethodPost:
			err := json.NewDecoder(req.Body).Decode(funcArgs)
			defer req.Body.Close()
			if err != nil {
				log.Panicf("error %v", err)
				return
			}

		case http.MethodGet:
			funcArgsTyp := reflect.TypeOf(funcArgs).Elem()
			for i := 0; i < funcArgsTyp.NumField(); i++ {
				jsonTag := funcArgsTyp.Field(i).Tag.Get("json")
				reflect.ValueOf(funcArgs).Elem().Field(i).Set(reflect.ValueOf(req.FormValue(jsonTag)))
			}
		}

		ctx := req.Context()
		outs := reflect.ValueOf(handlerFunc).Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(funcArgs)})
		resultFunc := outs[0].Interface()
		errFunc := outs[1].Interface()

		resp.Header().Set("Content-Type", "application/json")
		if errFunc != nil {
			err := errFunc.(xerror.XError)
			resp.WriteHeader(err.HttpStatus())
			json.NewEncoder(resp).Encode(map[string]string{
				"error": err.Message,
			})
			return
		} else {
			// json.NewEncoder(resp).Encode(map[string]interface{}{
			// 	"data": resultFunc,
			// })

			json.NewEncoder(resp).Encode(resultFunc)
		}

	case HTML:
		reflect.ValueOf(handlerFunc).Call([]reflect.Value{reflect.ValueOf(resp), reflect.ValueOf(req)})
	}
}

func (s *RemiService) validToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.jwtKey), nil
	})
	if err != nil {
		log.Println(err)
		return req, false
	}

	if !t.Valid {
		return req, false
	}

	id, ok := claims["id"].(string)
	if !ok {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), accountAuthKey(0), id))
	return req, true
}

type accountAuthKey int8

func accountIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(accountAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
