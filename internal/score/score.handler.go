package score

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler interface {
	Increment(c *gin.Context)
	GetAllScore(c *gin.Context)
}

func NewHandler(service Service, logger *zap.Logger) Handler {
	return &handlerImpl{
		service, logger,
	}
}

type handlerImpl struct {
	service Service
	logger  *zap.Logger
}

func (h *handlerImpl) Increment(c *gin.Context) {}

func (h *handlerImpl) GetAllScore(c *gin.Context) {}
