package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"social-network/internal/entities"
	"social-network/internal/repositories"
	"social-network/up"
	"social-network/utils/crypto"
	"social-network/utils/golibs/idutil"
	"social-network/utils/xerror"

	"github.com/dgrijalva/jwt-go"
)

var _ up.AccountService = &AccountService{}

type AccountService struct {
	accountRepo *repositories.AccountRepository
	jwtKey      string
}

func NewAccountService(db *sql.DB, jwtKey string) *AccountService {
	return &AccountService{
		accountRepo: repositories.NewAccountRepository(db),
		jwtKey:      jwtKey,
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
