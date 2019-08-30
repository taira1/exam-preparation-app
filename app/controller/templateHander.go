package controller

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

// TemplateHandler テンプレートハンドラ
type TemplateHandler struct {
	once       sync.Once
	Filename   string
	templ      *template.Template
	Controller controller
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//テンプレートのコンパイル
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.Filename)))
	})
	data := t.Controller.process(w, r)
	t.templ.Execute(w, data)

}
