package controller

import (
	"html/template"
	"net/http"
)

type controller interface {
	process(w http.ResponseWriter, r *http.Request) map[string]interface{}
	specifyTemplate() *template.Template //コントローラの各アクションに基づいて表示するhtmlのtemplateインスタンスを返す

}
