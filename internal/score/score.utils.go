package score

import (
	"github.com/bookpanda/angryfern-backend/internal/dto"
	"github.com/bookpanda/angryfern-backend/internal/model"
)

func (s *serviceImpl) modelToDTO(score *model.Score) dto.Score {
	return dto.Score{
		Code:        score.Code,
		CountryName: score.CountryName,
		ClickCount:  score.ClickCount,
	}
}

func (s *serviceImpl) modelToDTOList(scores *[]model.Score) []dto.Score {
	var results []dto.Score
	for _, score := range *scores {
		results = append(results, s.modelToDTO(&score))
	}
	return results
}
