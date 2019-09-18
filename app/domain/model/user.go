package model

// User ユーザモデルです。
type User struct {
	ID        int
	Name      string
	Comment   string
	Education *Subject
}
