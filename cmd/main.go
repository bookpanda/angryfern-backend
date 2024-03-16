package main

import (
	"fmt"

	"github.com/bookpanda/angryfern-backend/cfgldr"
	"github.com/bookpanda/angryfern-backend/database"
	healthcheck "github.com/bookpanda/angryfern-backend/internal/health_check"
	"github.com/bookpanda/angryfern-backend/internal/middleware"
	"github.com/bookpanda/angryfern-backend/internal/router"
	"github.com/bookpanda/angryfern-backend/internal/score"
	"github.com/bookpanda/angryfern-backend/logger"
)

func main() {
	conf, err := cfgldr.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	db, err := database.New(conf)
	if err != nil {
		panic(fmt.Sprintf("Failed to load database: %v", err))
	}

	logger := logger.New(conf)

	hcHandler := healthcheck.NewHandler()

	scoreRepo := score.NewRepository(db)
	scoreSvc := score.NewService(scoreRepo, logger)
	scoreHdr := score.NewHandler(scoreSvc, logger)

	corsHandler := cfgldr.MakeCorsConfig(conf)
	appMiddleware := middleware.NewAppMiddleware(conf)
	r := router.New(conf, corsHandler, appMiddleware)

	r.GET("/hc", hcHandler.HealthCheck)
	r.PUT("/scores/add", scoreHdr.Increment)
	r.GET("/scores", scoreHdr.GetAllScore)

	if err := r.Run(fmt.Sprintf(":%v", conf.AppConfig.Port)); err != nil {
		logger.Fatal("unable to start server")
	}
}
