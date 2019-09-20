package controller

import (
	"html/template"
	"net/http"

	"github.com/stretchr/objx"
)

type controller interface {
	process(w http.ResponseWriter, r *http.Request) map[string]interface{}
	specifyTemplate() *template.Template //コントローラの各アクションに基づいて表示するhtmlのtemplateインスタンスを返す
}

// RedirectTo 指定したURLへリダイレクトします。
func RedirectTo(w http.ResponseWriter, URL string) {
	w.Header().Set("Location", URL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

// setErrMessageToCookie エラーメッセージをクッキーにセットします。
func setErrMessageToCookie(w http.ResponseWriter, r *http.Request, errMessage string) {
	ErrCookieValue := objx.New(map[string]interface{}{
		"Messages": errMessage,
	}).MustBase64()
	if cookieValue, err := r.Cookie("err"); err == nil {
		cookieValue.Value = ErrCookieValue
		http.SetCookie(w, cookieValue)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "err",
		Value: ErrCookieValue,
		Path:  "/"})
}

// getErrMessagesFromCookie Cookieからエラーメッセージを取得します。
func getErrMessagesFromCookie(r *http.Request) interface{} {
	if cookieValue, err := r.Cookie("err"); err == nil {
		if cookieValue.Value == "" {
			return nil
		}
		if m, ok := objx.MustFromBase64(cookieValue.Value)["Messages"].(string); ok {
			return m
		}
	}
	return nil
}

// deleteCookieByName 指定した名前のCookieを削除します。
func deleteCookieByName(w http.ResponseWriter, r *http.Request, cookieName string) {
	if c, err := r.Cookie(cookieName); err == nil {
		c.Value = ""
		c.Path = ""
		http.SetCookie(w, c)
	}
}
