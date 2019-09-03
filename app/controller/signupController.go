package controller

import (
	"net/http"
)

// SignupController サインアップコントローラです。
type SignupController struct {
}

func (c *SignupController) process(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	return nil
}

// SigunpHandler サインアップハンドラ
func SigunpHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// email := r.Form["email"]
	// password := r.Form["password"]
	// university := r.Form["university"]
	//アカウント作成に必要な情報
	//email, password
	//user_name 学歴
	//
	//
	//
	//TODO: 入力内容が妥当かvalidationをかける
	//TODO:　妥当であればuser,authテーブルに必要情報をインサートする。

}
