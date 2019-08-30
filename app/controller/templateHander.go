package controller

import (
	"exam-preparation-app/app/trace"
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
	Tracer     trace.Tracer
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//テンプレートのコンパイル
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("../views/templates", t.Filename)))
	})
	data := t.Controller.process(w, r)
	t.templ.Execute(w, data)

}
