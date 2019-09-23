package controller

import (
	"net/http"

	"github.com/stretchr/objx"
)

// setErrMessageToCookie エラーメッセージをクッキーにセットします。
func setErrMessageToCookie(w http.ResponseWriter, r *http.Request, errMessage string) {
	ErrCookieValue := objx.New(map[string]interface{}{
		"Messages": errMessage,
	}).MustBase64()
	setValueToCookie(w, r, ErrCookieValue, "err")
}

// setNewErrMessageToCookie エラーメッセージをクッキーにセットします。
func setNewErrMessageToCookie(w http.ResponseWriter, r *http.Request, errMessage string) {
	ErrCookieValue := objx.New(map[string]interface{}{
		"Messages": errMessage,
	}).MustBase64()
	http.SetCookie(w, &http.Cookie{
		Name:  "err",
		Value: ErrCookieValue,
		Path:  "/"}) //TODO: Pathが"/"なのはセキュリティ上よくないので、厳密に指定する。

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

type confirmeInfo struct {
	NextURL       string
	Title         string
	Message       string
	ButtonMessage string
}

func setConfirmeInfoToCookie(w http.ResponseWriter, r *http.Request, value *confirmeInfo) {
	confirmeInfoCookieValue := objx.New(map[string]interface{}{
		"NextURL":       value.NextURL,
		"Title":         value.Title,
		"Message":       value.Message,
		"ButtonMessage": value.ButtonMessage,
	}).MustBase64()
	setValueToCookie(w, r, confirmeInfoCookieValue, "confirmeInfo")
}

// getConfirmeInfoFromCookie CookieからconfirmeInfoを取得します。
func getConfirmeInfoFromCookie(r *http.Request) *confirmeInfo {
	confirmeInfo := &confirmeInfo{}
	if cookieValue, err := r.Cookie("confirmeInfo"); err == nil {
		if cookieValue.Value == "" {
			return confirmeInfo
		}
		if t, ok := objx.MustFromBase64(cookieValue.Value)["Title"].(string); ok {
			confirmeInfo.Title = t
		}
		if n, ok := objx.MustFromBase64(cookieValue.Value)["NextURL"].(string); ok {
			confirmeInfo.NextURL = n
		}
		if m, ok := objx.MustFromBase64(cookieValue.Value)["Message"].(string); ok {
			confirmeInfo.Message = m
		}
		if b, ok := objx.MustFromBase64(cookieValue.Value)["ButtonMessage"].(string); ok {
			confirmeInfo.ButtonMessage = b
		}
		return confirmeInfo
	}
	return nil
}

func setValueToCookie(w http.ResponseWriter, r *http.Request, value string, name string) {
	if cookieValue, err := r.Cookie(name); err == nil {
		cookieValue.Value = value
		http.SetCookie(w, cookieValue)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  name,
		Value: value,
		Path:  "/"}) //TODO: Pathが"/"なのはセキュリティ上よくないので、厳密に指定する。
}
