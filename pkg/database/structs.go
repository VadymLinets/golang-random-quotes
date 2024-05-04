package database

type Quote struct {
	ID     string
	Quote  string
	Author string
	Likes  int64
}

type View struct {
	UserID  string
	QuoteID string
	Liked   bool
}
