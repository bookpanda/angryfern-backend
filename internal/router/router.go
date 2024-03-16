package router

import (
	"github.com/bookpanda/angryfern-backend/cfgldr"
	"github.com/bookpanda/angryfern-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func New(conf *cfgldr.Config, corsHandler cfgldr.CorsHandler, appMiddleware middleware.AppMidddleware) *Router {
	if !conf.AppConfig.IsDevelopment() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(gin.HandlerFunc(corsHandler))
	r.Use(gin.HandlerFunc(appMiddleware))

	return &Router{r}
}
