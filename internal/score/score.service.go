package score

import (
	"github.com/bookpanda/angryfern-backend/errors"
	"github.com/bookpanda/angryfern-backend/internal/dto"
	"github.com/bookpanda/angryfern-backend/internal/model"
	"go.uber.org/zap"
)

type Service interface {
	Increment(req *dto.IncrementScoreRequest) (*dto.IncrementScoreResponse, *errors.AppError)
	GetAllScore() (*dto.GetAllScoreResponse, *errors.AppError)
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

func (s *serviceImpl) Increment(req *dto.IncrementScoreRequest) (*dto.IncrementScoreResponse, *errors.AppError) {
	err := s.repo.Increment(req.Code, req.Amount)
	if err != nil {
		s.logger.Error("failed to increment score", zap.Error(err))
		return nil, errors.InternalError
	}

	return &dto.IncrementScoreResponse{Success: true}, nil
}

func (s *serviceImpl) GetAllScore() (*dto.GetAllScoreResponse, *errors.AppError) {
	var query []model.Score
	err := s.repo.GetAll(&query)
	if err != nil {
		s.logger.Error("failed to get all scores", zap.Error(err))
		return nil, errors.InternalError
	}

	results := s.modelToDTOList(&query)

	return &dto.GetAllScoreResponse{Scores: results}, nil
}
