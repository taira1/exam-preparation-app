package controller

import (
	"net/http"

	"github.com/stretchr/objx"
)

// UserController ユーザコントローラです。
type UserController struct {
}

func process(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	// ユーザ情報をCookieに記録
	if authCookie, err := r.Cookie("auth"); err == nil {
		return map[string]interface{}{
			"User": objx.MustFromBase64(authCookie.Value),
		}
	}
	return nil
}