package logger

import (
	"github.com/bookpanda/angryfern-backend/cfgldr"
	"go.uber.org/zap"
)

func InitLogger(conf *cfgldr.Config) *zap.Logger {
	var logger *zap.Logger

	if conf.AppConfig.IsDevelopment() {
		logger = zap.Must(zap.NewDevelopment())
	} else {
		logger = zap.Must(zap.NewProduction())
	}

	return logger
}
