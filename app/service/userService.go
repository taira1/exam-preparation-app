package service

import (
	"exam-preparation-app/app/domain/model"
	"exam-preparation-app/app/infrastructure"
)

// UserService ユーザサービスです
type UserService struct {
	validater *Validater
}

// ValidateUpdateUser User情報更新の妥当性を検証します
func (s *UserService) ValidateUpdateUser(name string, comment string) bool {
	if s.validater.ValidateNameLength(name) == false {
		return false
	}
	if s.validater.ValidateCommentLength(comment) == false {
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
