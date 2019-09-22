package controller

import (
	"html/template"
	"net/http"
)

// LogoutController ログアウトコントローラ
type LogoutController struct {
	htmlFilename string
}

func (c *LogoutController) process(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c.htmlFilename = "logout.html"
	return nil
}

func (c *LogoutController) specifyTemplate() *template.Template {
	return templateHelperOBJ.compiledTemplates[c.htmlFilename]
}

// LogoutHander ログアウトします。
func LogoutHander(w http.ResponseWriter, r *http.Request) {
	deleteCookieByName(w, r, "user")
	deleteCookieByName(w, r, "article")
	deleteCookieByName(w, r, "err")
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
