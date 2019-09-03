package service

import (
	"exam-preparation-app/app/infrastructure"
	"log"
	"regexp"
	"sync"
)

// Validater 検証機です
type Validater struct {
	once        sync.Once
	emailRegexp *regexp.Regexp
}

// ValidateEmail emailの妥当性をチェック
func (v *Validater) ValidateEmail(email string) bool {
	v.once.Do(func() {
		v.emailRegexp = regexp.MustCompile(`/[^\s]@[^\s]/`)
	})
	return v.emailRegexp.MatchString(email)
}

// ValidateUnique メールアドレスの一意性を検証します。
func (v *Validater) ValidateUnique(email string) bool {
	if infrastructure.InfrastructureOBJ.AuthAccesser.FindByEmail(email) != nil {
		log.Println("すでに登録済のメールアドレスです。")
		return false
	}
	return true
}
