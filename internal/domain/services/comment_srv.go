package services

import (
	"cook-book-backend/internal/domain/repositories"
	"cook-book-backend/internal/interfaces/dto"
)

type CommentService interface {
	GetCommentList(req *dto.GetCommentRequest) ([]dto.CommentResponse, int64, error)
	CreateComment(req *dto.CommentResponse) error
	CreateLike(commentId int64) error
	DeleteLike(commentId int64) error
}

type commentService struct {
	commentRepo repositories.CommentRepository
}

func NewCommentService(commentRepo repositories.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

// GetCommentList returns a list of comment based on the given request
func (cr *commentService) GetCommentList(req *dto.GetCommentRequest) ([]dto.CommentResponse, int64, error) {
	return cr.commentRepo.GetCommentList(req)
}

// CreateComment creates a new comment
func (cr *commentService) CreateComment(req *dto.CommentResponse) error {
	return cr.commentRepo.CreateComment(req)
}

// CreateLike creates a like
func (cr *commentService) CreateLike(commentId int64) error {
	return cr.commentRepo.CreateLike(commentId)
}

// DeleteLike deletes a like
func (cr *commentService) DeleteLike(commentId int64) error {
	return cr.commentRepo.DeleteLike(commentId)
}
