package model

// Auth 認証情報のモデルです
type Auth struct {
	ID       int
	Email    string
	Password string
	UserID   int
}
