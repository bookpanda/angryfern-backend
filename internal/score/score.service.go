package score

import "github.com/bookpanda/angryfern-backend/internal/dto"

type Service interface {
	Increment(req *dto.IncrementScoreRequest) (dto.IncrementScoreResponse, error)
	GetAllScore() (dto.GetAllScoreResponse, error)
}

func NewService(repo Repository) Service {
	return &serviceImpl{
		repo,
	}
}

type serviceImpl struct {
	repo Repository
}

func (s *serviceImpl) Increment(req *dto.IncrementScoreRequest) (dto.IncrementScoreResponse, error) {
	return dto.IncrementScoreResponse{}, nil
}

func (s *serviceImpl) GetAllScore() (dto.GetAllScoreResponse, error) {
	return dto.GetAllScoreResponse{}, nil
}
