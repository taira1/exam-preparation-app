package controller

import (
	"html/template"
	"net/http"
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
