package router

import (
	"github.com/bookpanda/angryfern-backend/cfgldr"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func New(conf *cfgldr.Config, corsHandler cfgldr.CorsHandler) *Router {
	if !conf.AppConfig.IsDevelopment() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(gin.HandlerFunc(corsHandler))

	return &Router{r}
}
