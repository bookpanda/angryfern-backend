package database

import (
	"github.com/bookpanda/angryfern-backend/cfgldr"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(conf *cfgldr.DatabaseConfig) (db *gorm.DB, err error) {
	return gorm.Open(postgres.Open(conf.Url), &gorm.Config{TranslateError: true})
}
