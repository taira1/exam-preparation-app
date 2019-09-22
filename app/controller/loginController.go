package controller

import (
	"exam-preparation-app/app/infrastructure"
	"exam-preparation-app/app/service"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
)

// LoginController ログインコントローラ
type LoginController struct {
	htmlFilename string
}

func (c *LoginController) process(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	defer deleteCookieByName(w, r, "err")
	c.htmlFilename = "login.html"
	return map[string]interface{}{
		"Err": getErrMessagesFromCookie(r),
	}
}

func (c *LoginController) specifyTemplate() *template.Template {
	return templateHelperOBJ.compiledTemplates[c.htmlFilename]
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
			setUserToCookie(w, user)
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
