package quote

import "quote/pkg/database"

const oneHundredPercent float64 = 100

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
