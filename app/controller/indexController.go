package controller

import (
	"html/template"
	"net/http"
)

// IndexController インデックスコントローラ
type IndexController struct {
	htmlFilename string
}

func (c *IndexController) process(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c.htmlFilename = "index.html"
	return nil
}

func (c *IndexController) specifyTemplate() *template.Template {
	return templateHelperOBJ.compiledTemplates[c.htmlFilename]
}
