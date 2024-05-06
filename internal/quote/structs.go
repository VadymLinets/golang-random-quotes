package quote

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

type Quote struct {
	ID     string   `json:"id"`
	Quote  string   `json:"quote"`
	Author string   `json:"author"`
	Tags   []string `json:"tags"`
	Likes  int64    `json:"likes"`
}

func fromDatabaseQuoteToQuote(quote database.Quote) Quote {
	return Quote{
		ID:     quote.ID,
		Quote:  quote.Quote,
		Author: quote.Author,
		Tags:   quote.Tags,
		Likes:  quote.Likes,
	}
}
