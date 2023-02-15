package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"social-network/internal/entities"
	"social-network/internal/repositories"
	"social-network/up"
	"social-network/utils/golibs/idutil"
	"social-network/utils/xerror"
)

var _ up.CommentService = &CommentService{}

type CommentService struct {
	commentRepo *repositories.CommentRepository
}

func NewCommentService(db *sql.DB) *CommentService {
	return &CommentService{
		commentRepo: repositories.NewCommentRepository(db),
	}
}

func (s *CommentService) Create(ctx context.Context, req *up.CreateCommentRequest) (*up.CreateCommentResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	currentAccount, _ := accountIDFromCtx(ctx)

	now := time.Now()
	comment := &entities.Comment{
		CommentID: idutil.NewID(),
		FeedID:    req.FeedID,
		AccountID: currentAccount,
		Message:   req.Message,
		ImageUrl:  req.ImageUrl,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	if err := s.commentRepo.Create(ctx, comment); err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.commentRepo.Create: %w", err))
	}

	return &up.CreateCommentResponse{
		CommentID: comment.CommentID,
	}, nil
}

func (s *CommentService) Update(ctx context.Context, req *up.UpdateCommentRequest) (*up.UpdateCommentResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	now := time.Now()
	comment := &entities.Comment{
		CommentID: req.CommentID,
		Message:   req.Message,
		ImageUrl:  req.ImageUrl,
		UpdatedAt: &now,
	}

	if err := s.commentRepo.Update(ctx, comment); err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.commentRepo.Update: %w", err))
	}

	return &up.UpdateCommentResponse{}, nil
}

func (s *CommentService) List(ctx context.Context, req *up.ListCommentsRequest) (*up.ListCommentsResponse, error) {
	listCommentsArgs := &repositories.ListCommentsArgs{
		FeedID: req.FeedID,
	}

	comments, err := s.commentRepo.List(ctx, listCommentsArgs)
	if err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.commentRepo.List: %w", err))
	}

	commentsUp := make([]*up.Comment, 0, len(comments))
	for _, comment := range comments {
		commentsUp = append(commentsUp, s.convertCommentEntToCommentUp(comment))
	}

	return &up.ListCommentsResponse{
		Comments: commentsUp,
	}, nil
}

func (s *CommentService) Delete(ctx context.Context, req *up.DeleteCommentRequest) (*up.DeleteCommentResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	currentAccount, _ := accountIDFromCtx(ctx)
	err := s.commentRepo.Delete(ctx, req.CommentID, currentAccount)
	if err != nil {
		return nil, xerror.Error(xerror.Internal, fmt.Errorf("s.commentRepo.Delete: %w", err))
	}

	return &up.DeleteCommentResponse{}, nil
}

func (s *CommentService) convertCommentEntToCommentUp(commentEnt *entities.Comment) *up.Comment {
	commentUp := &up.Comment{
		CommentID: commentEnt.CommentID,
		FeedID:    commentEnt.FeedID,
		AccountID: commentEnt.AccountID,
		Message:   commentEnt.Message,
		ImageUrl:  commentEnt.ImageUrl,
	}
	if commentEnt.CreatedAt != nil {
		commentUp.CreatedAt = *commentEnt.CreatedAt
	}
	if commentEnt.UpdatedAt != nil {
		commentUp.UpdatedAt = *commentEnt.UpdatedAt
	}

	return commentUp
}
