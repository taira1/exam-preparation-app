package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

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
			setErrMessageToCookie(w, r, "ログインが必要です")
			http.Redirect(w, r, fmt.Sprintf("/login"), http.StatusTemporaryRedirect)
			return nil
		}
		realUserID := objx.MustFromBase64(userCookie.Value)["userID"]
		userID, _ := strconv.Atoi(segs[2])
		if realUserID.(int) != userID {
			setErrMessageToCookie(w, r, "cookieの情報が壊れています")
			http.Redirect(w, r, fmt.Sprintf("/login"), http.StatusTemporaryRedirect)
			return nil
		}
	// case "article": //
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
