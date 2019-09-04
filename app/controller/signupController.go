package controller

import (
	"exam-preparation-app/app/infrastructure"
	"net/http"
)

// SignupController サインアップコントローラです。
type SignupController struct {
}

func (c *SignupController) process(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	//TODO: 大学、学科、学部情報をフロントに渡す
	universities := infrastructure.InfrastructureOBJ.UniversityAccesser.FindAll()
	faculties := infrastructure.InfrastructureOBJ.FacultyAccesser.FindAll()
	subjects := infrastructure.InfrastructureOBJ.SubjectAccesser.FindAll()
	data := map[string]interface{}{
		"Universities": universities,
		"Faculties":    faculties,
		"Subjects":     subjects,
	}
	return data
}

// SigunpHandler サインアップハンドラ
func SigunpHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
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
	return nil

}
