package model

import "exam-preparation-app/app/crypto"

// Auth 認証情報のモデルです
type Auth struct {
	ID       int
	Email    string
	Password string
	UserID   int
}

// NewAuthToBeRegisterd 登録予定のAuthインスタンスを作成します。パスワードはこの関数内でハッシュ化されます。
func NewAuthToBeRegisterd(email string, password string) (*Auth, error) {
	HashedPassword, err := crypto.PasswordEncrypt(password)
	return &Auth{
		Email:    email,
		Password: HashedPassword,
	}, err
}
