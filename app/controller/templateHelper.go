package controller

import (
	"html/template"
	"path/filepath"

	"github.com/playree/goingtpl"
)

var templateHelperOBJ = newTemplateHelper()

type templateHelper struct {
	compiledTemplates map[string]*template.Template
}

func newTemplateHelper() *templateHelper {
	t := &templateHelper{compiledTemplates: make(map[string]*template.Template)}
	t.setCompileTemplate("index.html")
	t.setCompileTemplate("login.html")
	t.setCompileTemplate("logout.html")
	t.setCompileTemplate("signup.html")
	t.setCompileTemplate("user_home.html")
	t.setCompileTemplate("user_edit.html")
	t.setCompileTemplate("user_article_list.html")
	t.setCompileTemplate("article_complete.html")
	t.setCompileTemplate("article_edit.html")
	t.setCompileTemplate("article_list.html")
	t.setCompileTemplate("complete.html")
	t.setCompileTemplate("error.html")
	t.setCompileTemplate("confirme.html")

	return t
}

func (t *templateHelper) setCompileTemplate(filename string) {
	t.compiledTemplates[filename] = template.Must(goingtpl.ParseFile(filepath.Join("../views/templates", filename)))
}
