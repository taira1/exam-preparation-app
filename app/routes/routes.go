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
	http.Handle("/signup", controller.NewTmplateHandler(&controller.SignupController{}))

	// TODO: ユーザの記事一覧などのハンドリングを考える。
	// http.Handler("/User/", &controller.TemplateHandler{
	// 	Filename: "User"
	// })
}
