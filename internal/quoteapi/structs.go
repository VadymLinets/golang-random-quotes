package quoteapi

import "quote/pkg/database"

type RandomQuote struct {
	ID      string   `json:"_id"`
	Content string   `json:"content"`
	Author  string   `json:"author"`
	Tags    []string `json:"tags"`
}

func (rq *RandomQuote) toDatabase() database.Quote {
	return database.Quote{
		ID:     rq.ID,
		Quote:  rq.Content,
		Author: rq.Author,
		Tags:   rq.Tags,
	}
}
