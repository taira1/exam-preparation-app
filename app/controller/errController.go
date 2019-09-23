package controller

import (
	"html/template"
	"net/http"
)

// ErrController 登録完了画面コントローラです。
type ErrController struct {
	htmlFilename string
}

func (c *ErrController) process(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c.htmlFilename = "error.html"
	return map[string]interface{}{
		"Err": getErrMessagesFromCookie(r),
	}
}

func (c *ErrController) specifyTemplate() *template.Template {
	return templateHelperOBJ.compiledTemplates[c.htmlFilename]
}
