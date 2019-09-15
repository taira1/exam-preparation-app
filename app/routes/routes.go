package routes

import (
	"exam-preparation-app/app/controller"
	"net/http"
)

// RegisterTheHandler 各種ハンドラをhttp.Handleに登録します。
func RegisterTheHandler() {
	http.HandleFunc("/auth/", controller.LoginHander)
	http.Handle("/login", controller.NewTemplateHandler(&controller.LoginController{}))
	http.HandleFunc("/signup/post", controller.SigunpHandler)
	http.Handle("/signup/complete", controller.NewTemplateHandler(&controller.CompleteController{}))
	http.Handle("/signup", controller.NewTemplateHandler(&controller.SignupController{}))
}
