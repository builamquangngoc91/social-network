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
	"social-network/utils/crypto"
	"social-network/utils/golibs/idutil"
	"social-network/utils/kafka"
	"social-network/utils/xerror"

	"github.com/dgrijalva/jwt-go"
)

var _ up.AccountService = &AccountService{}

type AccountService struct {
	accountRepo   *repositories.AccountRepository
	followerRepo  *repositories.FollowerRepository
	jwtKey        string
	kafkaProducer *kafka.KafkaProducer
}

func NewAccountService(db *sql.DB, jwtKey string, kafkaProducer *kafka.KafkaProducer) *AccountService {
	return &AccountService{
		accountRepo:   repositories.NewAccountRepository(db),
		followerRepo:  repositories.NewFollowerRepository(db),
		jwtKey:        jwtKey,
		kafkaProducer: kafkaProducer,
	}
}

func (s *AccountService) Register(ctx context.Context, req *up.RegisterRequest) (*up.RegisterResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, xerror.Error(xerror.InvalidArgument, err)
	}

	account, err := s.accountRepo.FindByUsername(ctx, req.Username)
	if err != nil && err != sql.ErrNoRows {
		return nil, xerror.Error(xerror.Internal, err)
	}

	if account != nil {
		return nil, xerror.Error(xerror.InvalidArgument, fmt.Errorf("account exists with the given username"))
	}

	password, err := crypto.HashPassword(req.Password)
	if err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("crypto.HashPassword: %w", err))
	}

	now := time.Now()
	err = s.accountRepo.Create(ctx, &entities.Account{
		AccountID: idutil.NewID(),
		Username:  req.Username,
		Password:  password,
		Fullname:  req.Fullname,
		Email:     req.Email,
		CreatedAt: &now,
		UpdatedAt: &now,
	})
	if err != nil {
		return nil, xerror.Error(xerror.Internal, err)
	}

	return &up.RegisterResponse{}, nil
}

func (s *AccountService) Login(ctx context.Context, req *up.LoginRequest) (*up.LoginResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	account, err := s.accountRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.accountRepo.FindByUsername: %w", err))
		}
		return nil, xerror.Error(xerror.UnAuthorized, fmt.Errorf("incorrect username/pwd"))
	}

	if account == nil || !crypto.CheckPasswordHash(req.Password, account.Password) {
		return nil, xerror.Error(xerror.UnAuthorized, fmt.Errorf("incorrect username/pwd"))
	}

	token, err := s.createToken(account.AccountID, account.Username)
	if err != nil {
		return nil, xerror.Error(xerror.Internal, err)
	}

	return &up.LoginResponse{
		ID:       account.AccountID,
		Username: account.Username,
		Fullname: account.Fullname,
		Email:    account.Email,
		Token:    token,
	}, nil
}

func (s *AccountService) createToken(id, username string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["id"] = id
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(s.jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AccountService) Follow(ctx context.Context, req *up.FollowRequest) (*up.FollowResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	currentAccount, _ := accountIDFromCtx(ctx)
	if _, err := s.accountRepo.FindByID(ctx, req.AccountID); err != nil {
		if err == sql.ErrNoRows {
			return nil, xerror.Error(xerror.Internal, fmt.Errorf("account %s not found", req.AccountID))
		}
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.accountRepo.FindByID: %w", err))
	}

	now := time.Now()
	follower := &entities.Follower{
		ID:           idutil.NewID(),
		AccountID:    currentAccount,
		FollowerID:   req.AccountID,
		FollowedDate: &now,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}

	if err := s.followerRepo.Upsert(ctx, follower); err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.followerRepo.Upsert: %w", err))
	}

	// send message to kafka
	body, err := handlers.NewAccountFollowed(currentAccount, req.AccountID)
	if err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("handlers.NewAccountFollowed: %w", err))
	}

	if err := s.kafkaProducer.SendMessage(ctx, string(kafkaConfig.EventTopic), body); err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.kafkaProducer.SendMessage: %w", err))
	}

	return &up.FollowResponse{}, nil
}

func (s *AccountService) UnFollow(ctx context.Context, req *up.UnFollowRequest) (*up.UnFollowResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	currentAccount, _ := accountIDFromCtx(ctx)
	if _, err := s.accountRepo.FindByID(ctx, req.AccountID); err != nil {
		if err == sql.ErrNoRows {
			return nil, xerror.Error(xerror.Internal, fmt.Errorf("account %s not found", req.AccountID))
		}
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.accountRepo.FindByID: %w", err))
	}

	now := time.Now()
	follower := &entities.Follower{
		ID:             idutil.NewID(),
		AccountID:      currentAccount,
		FollowerID:     req.AccountID,
		UnFollowedDate: &now,
		CreatedAt:      &now,
		UpdatedAt:      &now,
	}

	if err := s.followerRepo.Upsert(ctx, follower); err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.followerRepo.Upsert: %w", err))
	}

	return &up.UnFollowResponse{}, nil
}
