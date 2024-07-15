package graphql

import (
	"quote/internal/heartbeat"
	"quote/internal/quote"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	QuotesHandler    *quote.Service
	HeartbeatHandler *heartbeat.Service
}
