package routes

import (
	"exam-preparation-app/app/controller"
)

func registerTheHandler() {
	template := &controller.TemplateHandler{Filename: "login.html"}

	//http.Handle("/login", &TemplateHandler{Filename: "login.html", Controller: &loginController{}})
}
