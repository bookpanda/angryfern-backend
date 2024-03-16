package score

import (
	"github.com/bookpanda/angryfern-backend/internal/dto"
	"go.uber.org/zap"
)

type Service interface {
	Increment(req *dto.IncrementScoreRequest) (dto.IncrementScoreResponse, error)
	GetAllScore() (dto.GetAllScoreResponse, error)
}

func NewService(repo Repository, logger *zap.Logger) Service {
	return &serviceImpl{
		repo, logger,
	}
}

type serviceImpl struct {
	repo   Repository
	logger *zap.Logger
}

func (s *serviceImpl) Increment(req *dto.IncrementScoreRequest) (dto.IncrementScoreResponse, error) {
	return dto.IncrementScoreResponse{}, nil
}

func (s *serviceImpl) GetAllScore() (dto.GetAllScoreResponse, error) {
	return dto.GetAllScoreResponse{}, nil
}
