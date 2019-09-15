package controller

import (
	"html/template"
	"net/http"
)

// LoginController ログインコントローラ
type LoginController struct {
}

func (c *LoginController) process(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	return nil
}

func (c *LoginController) specifyTemplate() *template.Template {
	return templateHelperOBJ.compiledTemplates["login.html"]
}
