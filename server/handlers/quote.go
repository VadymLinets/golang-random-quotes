package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetQuoteHandler(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		slog.Error("GetQuoteHandler got empty user_id")
		c.Status(http.StatusBadRequest)
		return
	}

	resp, err := h.some.GetQuote(c, userID)
	if err != nil {
		slog.Error("Failed to get quote:", "err", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) LikeQuoteHandler(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		slog.Error("LikeQuoteHandler got empty user_id")
		c.Status(http.StatusBadRequest)
		return
	}

	quoteID := c.Query("quote_id")
	if quoteID == "" {
		slog.Error("LikeQuoteHandler got empty quote_id")
		c.Status(http.StatusBadRequest)
		return
	}

	err := h.some.LikeQuote(c, userID, quoteID)
	if err != nil {
		slog.Error("Failed to like quote:", "quote id", quoteID, "err", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetSameQuoteHandler(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		slog.Error("GetSameQuoteHandler got empty user_id")
		c.Status(http.StatusBadRequest)
		return
	}

	quoteID := c.Query("quote_id")
	if quoteID == "" {
		slog.Error("GetSameQuoteHandler got empty quote_id")
		c.Status(http.StatusBadRequest)
		return
	}

	quote, err := h.some.GetSameQuote(c, userID, quoteID)
	if err != nil {
		slog.Error("Failed to get same quote:", "quote id", quoteID, "err", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, quote)
}
