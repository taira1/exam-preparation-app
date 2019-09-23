package service

import (
	"database/sql"
	"exam-preparation-app/app/crypto"
	"exam-preparation-app/app/domain/model"
	"exam-preparation-app/app/infrastructure"
	"exam-preparation-app/app/infrastructure/dataAccess"
)

// AuthService 認証関係の処理を担当します。
type AuthService struct {
	Validater *Validater
}

// Authenticate 指定したemailアドレスが認証できた場合そのユーザのidを返します。認証に失敗した場合-1を返します。
func Authenticate(email string, pass string) int {
	authObj := infrastructure.InfrastructureOBJ.AuthAccesser.FindByEmail(email)
	if authObj == nil {
		return -1
	}
	if crypto.CompareHashAndPassword(authObj.Password, pass) {
		return authObj.UserID
	}
	return -1
}

// ValidateAuth 認証情報の妥当性を検証します
func (s *AuthService) ValidateAuth(email string) bool {
	if s.Validater.ValidateEmailUnique(email) == false {
		return false
	}
	if s.Validater.ValidateEmail(email) == false {
		return false
	}
	return true
}

// RegisterAuthUser 引数で受け取ったauth,Userを登録します。
func RegisterAuthUser(auth *model.Auth, user *model.User) {
	dataAccess.Transact(infrastructure.InfrastructureOBJ.DBAgent.Conn, func(tx *sql.Tx) error {
		var err error
		userID := infrastructure.InfrastructureOBJ.UserAccesser.Insert(user)
		if userID == -1 {
			return err
		}
		infrastructure.InfrastructureOBJ.AuthAccesser.Insert(auth, userID)
		return err
	})
}
