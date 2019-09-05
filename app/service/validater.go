package service

import (
	"exam-preparation-app/app/infrastructure"
	"log"
	"regexp"
	"sync"
	"unicode/utf8"
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

// ValidateEmailUnique メールアドレスの一意性を検証します。
func (v *Validater) ValidateEmailUnique(email string) bool {
	if infrastructure.InfrastructureOBJ.AuthAccesser.FindByEmail(email) != nil {
		log.Println("すでに登録済のメールアドレスです。")
		return false
	}
	return true
}

// ValidatePasswordLength パスワードの文字数を検証します。
func (v *Validater) ValidatePasswordLength(password string) bool {
	return v.validateStringLength(password, 8, 70)
}

// ValidateNameLength ニックネームの文字数を検証します
func (v *Validater) ValidateNameLength(name string) bool {
	return v.validateStringLength(name, 1, 30)
}

func (v *Validater) validateStringLength(str string, max int, min int) bool {
	if min <= utf8.RuneCountInString(str) && utf8.RuneCountInString(str) <= max {
		return true
	}
	return false
}

// ValidateExistenceOfSubject 指定したIDの学科の存在性を検証します
func (v *Validater) ValidateExistenceOfSubject(subjectID int) bool {
	if infrastructure.InfrastructureOBJ.SubjectAccesser.FindByID(subjectID) != nil {
		return true
	}
	return false
}
