package database

import (
	"log"

	"github.com/bookpanda/angryfern-backend/cfgldr"
	"github.com/bookpanda/angryfern-backend/internal/model"
	"github.com/bookpanda/angryfern-backend/internal/model/seeds"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(conf *cfgldr.DatabaseConfig) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(conf.Url), &gorm.Config{TranslateError: true})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Score{})
	if err != nil {
		return nil, err
	}

	var count int64
	for _, b := range seeds.Countries {
		if db.Model(&model.Score{}).Where("code = ?", b.Code).Count(&count); count == 0 {
			err := db.Create(&b).Error
			if err != nil {
				return nil, err
			}
		} else {
			err := db.Where("code = ?", b.Code).Updates(&b).Error
			if err != nil {
				return nil, err
			}
		}
	}
	log.Println("✔️Seed", "Countries", "succeed")

	return
}
