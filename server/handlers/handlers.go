package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"quote/internal/heartbeat"
	"quote/internal/quote"
)

type Handler struct {
	quotes    *quote.Service
	heartbeat *heartbeat.Service
}

func (h *Handler) Heartbeat(c *gin.Context) {
	err := h.heartbeat.Ping(c)
	if err != nil {
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
