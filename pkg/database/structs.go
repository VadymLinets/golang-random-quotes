package database

import "github.com/lib/pq"

type Quote struct {
	ID     string
	Quote  string
	Author string
	Tags   pq.StringArray `gorm:"type:text[]"`
	Likes  int64
}

type View struct {
	UserID  string
	QuoteID string
	Liked   bool
}
