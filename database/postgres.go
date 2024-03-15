package database

import (
	"bookpanda/angryfern-backend/cfgldr"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(conf *cfgldr.DatabaseConfig, isDebug bool) (db *gorm.DB, err error) {
	return gorm.Open(postgres.Open(conf.Url), &gorm.Config{TranslateError: true})
}
