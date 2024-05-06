package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"quote/internal/heartbeat"
	"quote/internal/quote"
)

type Handler struct {
	quotes    *quote.Service
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

func NewHandler(quotes *quote.Service, heartbeat *heartbeat.Service) *Handler {
	return &Handler{
		quotes:    quotes,
		heartbeat: heartbeat,
	}
}
