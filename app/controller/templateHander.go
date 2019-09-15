package controller

import (
	"exam-preparation-app/app/trace"
	"html/template"
	"net/http"
	"os"
)

// TemplateHandler テンプレートハンドラ
type TemplateHandler struct {
	templ      *template.Template
	Controller controller
	Tracer     trace.Tracer
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Host": r.Host,
	}
	data = t.Controller.process(w, r)
	t.templ = t.Controller.specifyTemplate()
	t.templ.Execute(w, data)
}

// NewTemplateHandler コンストラクタです
func NewTemplateHandler(c controller) *TemplateHandler {
	return &TemplateHandler{
		Controller: c,
		Tracer:     trace.New(os.Stdout),
	}
}
