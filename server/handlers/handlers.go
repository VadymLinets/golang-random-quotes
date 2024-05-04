package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"quote/internal/heartbeat"
	"quote/internal/quote"
)

type Handler struct {
	some      *quote.Service
	heartbeat *heartbeat.Service
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
	}
}
