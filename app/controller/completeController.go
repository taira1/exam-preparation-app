package controller

import (
	"html/template"
	"net/http"
)

// CompleteController 登録完了画面コントローラです。
type CompleteController struct {
	htmlFilename string
}

func (c *CompleteController) process(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c.htmlFilename = "complete.html"
	return nil
}

func (c *CompleteController) specifyTemplate() *template.Template {
	return templateHelperOBJ.compiledTemplates[c.htmlFilename]
}
