package model

// Faculty 学部モデルです。
type Faculty struct {
	ID         int
	Name       string
	University *University
}
