package handler

import (
	"cosmosdb-gin/pkg/apperror"
	"cosmosdb-gin/pkg/logger"
	"cosmosdb-gin/pkg/usecase"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type Handler struct {
	logger.ILogger
	uc usecase.IUseCase
}

func NewHandler(l logger.ILogger, uc usecase.IUseCase) *Handler {
	return &Handler{l, uc}
}

func HandleError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	switch {
	case errors.Is(err, apperror.NotFound):
		c.AbortWithStatus(404)

	default:
		c.AbortWithError(500, err)
	}
	return true
}

func (h *Handler) GetFamily(c *gin.Context) {
	id := c.Param("id")
	h.Info("get family", id)

	row, err := h.uc.GetFamily(c, c.Param("id"))

	if HandleError(c, err) {
		return
	}

	c.JSON(200, gin.H{"result": row})
}

func (h *Handler) Hello(c *gin.Context) {
	c.String(200, "hello")
}

var Wired = wire.NewSet(NewHandler, usecase.Wired)
