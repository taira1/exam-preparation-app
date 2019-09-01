package model

// User ユーザモデルです。
type User struct {
	ID        int
	Name      string
	Education *Subject
}
