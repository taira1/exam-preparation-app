package model

import "time"

// Article 記事モデルです
type Article struct {
	ID         int
	UserID     int
	LastUpdate time.Time
	Title      string
	Class      string
	Teacher    string
	Content    string
	Status     string
}
