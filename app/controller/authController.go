package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
)

type authController struct {
	next controller
}

func (h *authController) process(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	_, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return nil
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return h.next.process(w, r)
}

func mustAuth(n controller) *authController {
	return &authController{next: n}
}

// LoginHander ログインハンドラです
func LoginHander(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		switch provider {
		case "default":
			//TODO:認証ロジックの記述
			r.ParseForm()
			fmt.Println("email: ", r.Form["email"])
			fmt.Println("pass: ", r.Form["password"])
		default:
			provider, err := gomniauth.Provider(provider)
			if err != nil {
				http.Error(w, fmt.Sprintf("認証プロバイダの取得に失敗しました。 %s: %s", provider, err), http.StatusBadRequest)
				return
			}
			loginURL, err := provider.GetBeginAuthURL(nil, nil)
			if err != nil {
				http.Error(w, fmt.Sprintf("GetBeginAuthURLの呼び出し中にエラーが発生しました。 %s: %s", provider, err), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Location", loginURL)
			w.WriteHeader(http.StatusTemporaryRedirect)
		}
	case "callback":
		switch provider {
		case "default":
			//TODO:認証ロジックからクッキー生成ロジックを追加
		default:
			provider, err := gomniauth.Provider(provider)
			if err != nil {
				http.Error(w, fmt.Sprintf("認証プロバイダの取得に失敗しました。%s: %s", provider, err), http.StatusBadRequest)
				return
			}

			creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
			if err != nil {
				http.Error(w, fmt.Sprintf("認証を完了できませんでした。%s: %s", provider, err), http.StatusInternalServerError)
				return
			}

			user, err := provider.GetUser(creds)
			if err != nil {
				http.Error(w, fmt.Sprintf("ユーザーの取得に失敗しました。 %s: %s", provider, err), http.StatusInternalServerError)
				return
			}

			// save some data
			authCookieValue := objx.New(map[string]interface{}{
				"name": user.Name(),
			}).MustBase64()
			http.SetCookie(w, &http.Cookie{
				Name:  "auth",
				Value: authCookieValue,
				Path:  "/"})

			w.Header().Set("Location", "/chat")
			w.WriteHeader(http.StatusTemporaryRedirect)
		}
	}
}
