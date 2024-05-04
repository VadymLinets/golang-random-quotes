package quote

type RandomQuote struct {
	ID      string `json:"_id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Quote struct {
	ID     string `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}
