package middleware

import (
	"strings"

	"github.com/bookpanda/angryfern-backend/cfgldr"
	"github.com/bookpanda/angryfern-backend/errors"
	"github.com/gin-gonic/gin"
)

type AppMidddleware gin.HandlerFunc

func NewAppMiddleware(conf *cfgldr.Config) AppMidddleware {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			errors.ResponseError(c, errors.Unauthorized)
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			errors.ResponseError(c, errors.InvalidToken)
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != conf.AppConfig.ApiKey {
			errors.ResponseError(c, errors.InvalidToken)
			c.Abort()
			return
		}

		c.Next()
	}
}
