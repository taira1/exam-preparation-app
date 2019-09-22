package routes

import (
	"exam-preparation-app/app/controller"
	"net/http"
)

// RegisterTheHandler 各種ハンドラをhttp.Handleに登録します。
func RegisterTheHandler() {
	http.Handle("/", controller.NewTemplateHandler(&controller.IndexController{}))
	http.HandleFunc("/auth/", controller.LoginHander)
	http.Handle("/login", controller.NewTemplateHandler(&controller.LoginController{}))
	http.Handle("/logout", controller.NewTemplateHandler(&controller.LogoutController{}))
	http.HandleFunc("/logout/post", controller.LogoutHander)
	http.HandleFunc("/signup/post", controller.SigunpHandler)
	http.Handle("/signup/complete", controller.NewTemplateHandler(&controller.CompleteController{}))
	http.Handle("/signup", controller.NewTemplateHandler(&controller.SignupController{}))
	http.Handle("/user/", controller.NewTemplateHandler(controller.MustAuth(&controller.UserController{})))
}
