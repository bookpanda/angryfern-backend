package score

import (
	"net/http"

	"github.com/bookpanda/angryfern-backend/errors"
	"github.com/bookpanda/angryfern-backend/internal/dto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler interface {
	Increment(c *gin.Context)
	GetAllScore(c *gin.Context)
}

func NewHandler(svc Service, logger *zap.Logger) Handler {
	return &handlerImpl{
		svc, logger,
	}
}

type handlerImpl struct {
	svc    Service
	logger *zap.Logger
}

func (h *handlerImpl) Increment(c *gin.Context) {
	var body dto.IncrementScoreRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		h.logger.Error("failed to bind json", zap.Error(err))
		errors.ResponseError(c, errors.BadRequest)
	}

	resp, apperr := h.svc.Increment(&body)
	if apperr != nil {
		errors.ResponseError(c, apperr)
	}

	c.JSON(http.StatusOK, resp)
}

func (h *handlerImpl) GetAllScore(c *gin.Context) {
	resp, apperr := h.svc.GetAllScore()
	if apperr != nil {
		errors.ResponseError(c, apperr)
	}

	c.JSON(http.StatusOK, resp)
}
