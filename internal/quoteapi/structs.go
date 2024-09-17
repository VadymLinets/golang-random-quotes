package quoteapi

import (
	"github.com/spf13/cast"

	"quote/pkg/database"
)

type RandomQuote struct {
	ID      int      `json:"id"`
	Content string   `json:"quote"`
	Author  string   `json:"author"`
	Tags    []string `json:"tags"`
}

func (rq *RandomQuote) toDatabase() database.Quote {
	return database.Quote{
		ID:     cast.ToString(rq.ID),
		Quote:  rq.Content,
		Author: rq.Author,
		Tags:   rq.Tags,
	}
}
