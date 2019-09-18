package service

import (
	"exam-preparation-app/app/domain/model"
	"exam-preparation-app/app/infrastructure"
	"log"
)

// いらない可能性が高い。

// UserService ユーザサービスです
type UserService struct {
	validater *Validater
}

// ValidateUpdateUser User情報更新の妥当性を検証します
func (s *UserService) ValidateUpdateUser(name string, comment string) bool {
	if s.validater.ValidateNameLength(name) == false {
		log.Fatal("ニックネームは1文字以上30文字以内にしてください")
		return false
	}
	if s.validater.ValidateCommentLength(comment) == false {
		log.Fatal("自己紹介は200文字以内にしてください")
		return false
	}
	return true
}

// UpdateUser User情報を更新します。
func (s *UserService) UpdateUser(user *model.User) bool {
	return infrastructure.InfrastructureOBJ.UserAccesser.Update(user)
}

// NewUserService コンストラクタです
func NewUserService() *UserService {
	return &UserService{validater: NewValidater()}
}
