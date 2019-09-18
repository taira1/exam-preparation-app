package controller

import (
	"html/template"
	"path/filepath"
)

var templateHelperOBJ = newTemplateHelper()

type templateHelper struct {
	compiledTemplates map[string]*template.Template
}

func newTemplateHelper() *templateHelper {
	t := &templateHelper{compiledTemplates: make(map[string]*template.Template)}
	t.setCompileTemplate("index.html")
	t.setCompileTemplate("login.html")
	t.setCompileTemplate("signup.html")
	t.setCompileTemplate("user_home.html")
	t.setCompileTemplate("user_edit.html")
	t.setCompileTemplate("complete.html")
	t.setCompileTemplate("error.html")
	return t
}

func (t *templateHelper) setCompileTemplate(filename string) {
	t.compiledTemplates[filename] = template.Must(template.ParseFiles(filepath.Join("../views/templates", filename)))
}
