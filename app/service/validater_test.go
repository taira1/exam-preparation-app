package service

import (
	"strings"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	v := NewValidater()
	email := "testemail"
	if v.ValidateEmail(email) == true {
		t.Fatalf("ValidateEmailが正しく動作していません:不適切なメールアドレスにtrueを返してます。")
	}

	email = "user@example.com"
	if v.ValidateEmail(email) == false {
		t.Fatalf("ValidateEmailが正しく動作していません:適切なメールアドレスにfalseを返してます。")
	}

}

func TestValidateNameLength(t *testing.T) {
	v := NewValidater()
	name := createStringOfSpecifiedLength(30)
	if v.ValidateNameLength(name) == false {
		t.Fatalf("ValidateNameLength()が正しく動作してません:30文字でfalseを返している")
	}

	name = createStringOfSpecifiedLength(31)
	if v.ValidateNameLength(name) == true {
		t.Fatalf("ValidateNameLength()が正しく動作してません:30文字でtrueを返している")
	}

	name = createStringOfSpecifiedLength(1)
	if v.ValidateNameLength(name) == false {
		t.Fatalf("ValidateNameLength()が正しく動作してません:1文字でfalseを返している")
	}

	name = createStringOfSpecifiedLength(31)
	if v.ValidateNameLength(name) == true {
		t.Fatalf("ValidateNameLength()が正しく動作してません:0文字でtrueを返している")
	}

}

func createStringOfSpecifiedLength(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteString("a")
	}
	return sb.String()
}
