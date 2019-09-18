package controller

import (
	"exam-preparation-app/app/domain/model"
	"exam-preparation-app/app/infrastructure"
	"exam-preparation-app/app/service"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// SignupController サインアップコントローラです。
type SignupController struct {
	htmlFilename string
}

func (c *SignupController) process(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	universities := infrastructure.InfrastructureOBJ.UniversityAccesser.FindAll()
	faculties := infrastructure.InfrastructureOBJ.FacultyAccesser.FindAll()
	subjects := infrastructure.InfrastructureOBJ.SubjectAccesser.FindAll()
	data := map[string]interface{}{
		"Universities": universities,
		"Faculties":    faculties,
		"Subjects":     subjects,
	}
	c.htmlFilename = "signup.html"
	return data
}

func (c *SignupController) specifyTemplate() *template.Template {
	return templateHelperOBJ.compiledTemplates[c.htmlFilename]
}

// SigunpHandler サインアップハンドラ
func SigunpHandler(w http.ResponseWriter, r *http.Request) {
	v := service.NewValidater()
	r.ParseForm()
	userName := r.FormValue("userName")
	email := r.FormValue("email")
	password := r.FormValue("password")
	subjectID, err := strconv.Atoi(r.FormValue("subject"))
	if !v.ValidateNameLength(userName) {
		http.Error(w, "ニックネームは30文字以内にしてください", http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, "存在しない学校、学部、学科です。", http.StatusInternalServerError)
		return
	}
	if !(v.ValidateEmail(email) && v.ValidateEmailUnique(email)) {
		http.Error(w, "emailアドレスが正しくないか、既に存在するアドレスです。", http.StatusInternalServerError)
		return
	}
	if !v.ValidatePasswordLength(password) {
		http.Error(w, "パスワードは8文字以上70文字以内に設定してください", http.StatusInternalServerError)
		return
	}
	if !v.ValidateExistenceOfSubject(subjectID) {
		http.Error(w, "存在しない学校、学部、学科です。", http.StatusInternalServerError)
		return
	}

	auth, err := model.NewAuthToBeRegisterd(email, password)
	if err != nil {
		http.Error(w, "認証情報登録に失敗しました", http.StatusInternalServerError)
		return
	}
	user := &model.User{
		Name:      userName,
		Comment:   "",
		Education: infrastructure.InfrastructureOBJ.SubjectAccesser.FindByID(subjectID),
	}

	service.RegisterAuthUser(auth, user)

	nextURL := fmt.Sprintf("/signup/complete")
	RedirectTo(w, nextURL)
}
