package model

// Subject 学科モデルです。
type Subject struct {
	ID      int
	Name    string
	Faculty *Faculty
}
