package controller

import (
	"html/template"
	"net/http"
)

// LoginController ログインコントローラ
type LoginController struct {
	htmlFilename string
}

func (c *LoginController) process(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c.htmlFilename = "login.html"
	return nil
}

func (c *LoginController) specifyTemplate() *template.Template {
	return templateHelperOBJ.compiledTemplates[c.htmlFilename]
}
