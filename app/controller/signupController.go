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
	defer deleteCookieByName(w, r, "err")
	c.htmlFilename = "signup.html"
	universities := infrastructure.InfrastructureOBJ.UniversityAccesser.FindAll()
	faculties := infrastructure.InfrastructureOBJ.FacultyAccesser.FindAll()
	subjects := infrastructure.InfrastructureOBJ.SubjectAccesser.FindAll()
	data := map[string]interface{}{
		"Universities": universities,
		"Faculties":    faculties,
		"Subjects":     subjects,
		"Err":          getErrMessagesFromCookie(r),
	}
	return data
}

func (c *SignupController) specifyTemplate() *template.Template {
	return templateHelperOBJ.compiledTemplates[c.htmlFilename]
}

// SigunpHandler サインアップハンドラ
func SigunpHandler(w http.ResponseWriter, r *http.Request) {
	deleteCookieByName(w, r, "err")
	v := service.NewValidater()
	r.ParseForm()
	userName := r.FormValue("userName")
	email := r.FormValue("email")
	password := r.FormValue("password")
	subjectID, err := strconv.Atoi(r.FormValue("subject"))
	if !v.ValidateNameLength(userName) {
		setErrMessageToCookie(w, r, "ニックネームは30文字以内にしてください")
		http.Redirect(w, r, "/signup", http.StatusTemporaryRedirect)
		return
	}
	if err != nil {
		setErrMessageToCookie(w, r, "存在しない学校、学部、学科です。")
		http.Redirect(w, r, "/signup", http.StatusTemporaryRedirect)
		return
	}
	if !(v.ValidateEmail(email) && v.ValidateEmailUnique(email)) {
		setErrMessageToCookie(w, r, "emailアドレスが正しくないか、既に存在するアドレスです。")
		http.Redirect(w, r, "/signup", http.StatusTemporaryRedirect)
		return
	}
	if !v.ValidatePasswordLength(password) {
		setErrMessageToCookie(w, r, "パスワードは8文字以上70文字以内に設定してください")
		http.Redirect(w, r, "/signup", http.StatusTemporaryRedirect)
		return
	}
	if !v.ValidateExistenceOfSubject(subjectID) {
		setErrMessageToCookie(w, r, "存在しない学校、学部、学科です。")
		http.Redirect(w, r, "/signup", http.StatusTemporaryRedirect)
		return
	}

	auth, err := model.NewAuthToBeRegisterd(email, password)
	if err != nil {
		setErrMessageToCookie(w, r, "認証情報登録に失敗しました")
		http.Redirect(w, r, "/signup", http.StatusTemporaryRedirect)
		return
	}
	user := &model.User{
		Name:      userName,
		Comment:   "",
		Education: infrastructure.InfrastructureOBJ.SubjectAccesser.FindByID(subjectID),
	}
	service.RegisterAuthUser(auth, user)
	http.Redirect(w, r, fmt.Sprintf("/signup/complete"), http.StatusTemporaryRedirect)
}
