package crypto

import "golang.org/x/crypto/bcrypt"

// PasswordEncrypt パスワードのハッシュ化を行います
func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// CompareHashAndPassword パスワードとハッシュが適合しているか確認します
func CompareHashAndPassword(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
