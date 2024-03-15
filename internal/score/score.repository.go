package score

import (
	"github.com/bookpanda/angryfern-backend/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]model.Score, error)
	Increment(code string, amount uint) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db,
	}
}

type repositoryImpl struct {
	db *gorm.DB
}

func (r *repositoryImpl) Increment(code string, amount uint) error {
	return r.db.Model(&model.Score{}).Where("code = ?", code).Update("count", gorm.Expr("count + ?", amount)).Error
}

func (r *repositoryImpl) GetAll() ([]model.Score, error) {
	var scores []model.Score
	err := r.db.Find(&scores).Error
	return scores, err
}
