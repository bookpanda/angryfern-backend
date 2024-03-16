package score

import (
	"github.com/bookpanda/angryfern-backend/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(result *[]model.Score) error
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
	return r.db.Model(&model.Score{}).Where("code = ?", code).Update("click_count", gorm.Expr("click_count + ?", amount)).Error
}

func (r *repositoryImpl) GetAll(result *[]model.Score) error {
	return r.db.Find(&result).Error
}
