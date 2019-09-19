package controller

import (
	"exam-preparation-app/app/infrastructure"
	"exam-preparation-app/app/service"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
)

// AuthController 認証コントローラです
type AuthController struct {
	next controller
}

func (h *AuthController) process(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	segs := strings.Split(r.URL.Path, "/")
	action := segs[1]
	switch action {
	case "user":
		userCookie, err := r.Cookie("user")
		if err == http.ErrNoCookie {
			http.Redirect(w, r, fmt.Sprintf("/login"), http.StatusTemporaryRedirect)
			return nil
		}
		realUserID := objx.MustFromBase64(userCookie.Value)["userID"]
		userID, _ := strconv.Atoi(segs[2])
		if realUserID.(int) != userID {
			log.Fatal("cookieの情報が壊れています")
			http.Redirect(w, r, fmt.Sprintf("/login"), http.StatusTemporaryRedirect)
			return nil
		}
	case "article":
		//TODO:ロジックを追加
	case "chat":
		userCookie, _ := r.Cookie("user")
		authCookie, _ := r.Cookie("auth")
		if userCookie == nil && authCookie == nil {
			log.Fatal("ログインしていません")
			http.Redirect(w, r, fmt.Sprintf("/login"), http.StatusTemporaryRedirect)
			return nil
		}
	case "admin":
	//TODO:ロジックを追加
	default:
		//TODO:エラーページへリダイレクト
	}
	return h.next.process(w, r)
}

func (h *AuthController) specifyTemplate() *template.Template {
	return h.next.specifyTemplate()
}

// MustAuth コンストラクタです
func MustAuth(n controller) *AuthController {
	return &AuthController{next: n}
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
			r.ParseForm()
			userID := service.Authenticate(r.FormValue("email"), r.FormValue("password"))
			if userID == -1 {
				http.Error(w, fmt.Sprintf("認証を完了できませんでした。"), http.StatusInternalServerError)
				return //TODO: エラーページへリダイレクトする
			}
			user := infrastructure.InfrastructureOBJ.UserAccesser.FindByID(userID)
			SetUserToCookie(w, "user", user)
			http.Redirect(w, r, fmt.Sprintf("/user/%d/profile", userID), http.StatusTemporaryRedirect)
		default:
			provider, err := gomniauth.Provider(provider)
			if err != nil {
				http.Error(w, fmt.Sprintf("認証プロバイダの取得に失敗しました。 %s: %s", provider, err), http.StatusBadRequest)
				return //TODO: エラーページへリダイレクトする
			}
			loginURL, err := provider.GetBeginAuthURL(nil, nil)
			if err != nil {
				http.Error(w, fmt.Sprintf("GetBeginAuthURLの呼び出し中にエラーが発生しました。 %s: %s", provider, err), http.StatusInternalServerError)
				return //TODO: エラーページへリダイレクトする
			}
			http.Redirect(w, r, loginURL, http.StatusTemporaryRedirect)
		}
	case "callback":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(w, fmt.Sprintf("認証プロバイダの取得に失敗しました。%s: %s", provider, err), http.StatusBadRequest)
			return //TODO: エラーページへリダイレクトする
		}

		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			http.Error(w, fmt.Sprintf("認証を完了できませんでした。%s: %s", provider, err), http.StatusInternalServerError)
			return //TODO: エラーページへリダイレクトする
		}

		user, err := provider.GetUser(creds)
		if err != nil {
			http.Error(w, fmt.Sprintf("ユーザーの取得に失敗しました。 %s: %s", provider, err), http.StatusInternalServerError)
			return //TODO: エラーページへリダイレクトする
		}

		// save some data
		authCookieValue := objx.New(map[string]interface{}{
			"name": user.Name(),
		}).MustBase64()
		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookieValue,
			Path:  "/"}) //TODO: Pathが"/"なのはセキュリティ上よくないので、厳密に指定する。
		http.Redirect(w, r, "/chat", http.StatusTemporaryRedirect)
	}
}
