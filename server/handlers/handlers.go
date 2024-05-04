package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"

	"quote/internal/heartbeat"
	"quote/internal/quote"
)

type Handler struct {
	some      *quote.Service
	heartbeat *heartbeat.Service

	validate *validator.Validate
	resty    *resty.Client
}

func (h *Handler) HeartBeat(c *gin.Context) {
	err := h.heartbeat.PingDatabase(c)
	if err != nil {
		slog.Error("Failed to ping database", "err", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func NewHandler(some *quote.Service, heartbeat *heartbeat.Service) *Handler {
	return &Handler{
		some:      some,
		heartbeat: heartbeat,
		resty:     resty.New(),
		validate:  validator.New(),
	}
}
