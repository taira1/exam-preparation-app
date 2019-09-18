package controller

import (
	"exam-preparation-app/app/domain/model"
	"exam-preparation-app/app/infrastructure"
	"exam-preparation-app/app/service"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/stretchr/objx"
)

// UserController ユーザコントローラです。
type UserController struct {
	htmlFilename string
}

func (c *UserController) process(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	segs := strings.Split(r.URL.Path, "/")
	userID, _ := strconv.Atoi(segs[2])
	userAction := segs[3]
	var method string
	if len(segs) == 5 {
		method = segs[4]
	}
	switch userAction {
	case "profile":
		switch method {
		case "":
			c.htmlFilename = "user_home.html"
			return map[string]interface{}{
				"User": infrastructure.InfrastructureOBJ.UserAccesser.FindByID(userID),
			}
		case "edit":
			c.htmlFilename = "user_edit.html" //TODO:適切なHTMLテンプレートに変更する。
			return map[string]interface{}{
				"User": infrastructure.InfrastructureOBJ.UserAccesser.FindByID(userID),
			}
		case "post":
			us := service.NewUserService()
			if us.ValidateUpdateUser(r.FormValue("userName"), r.FormValue("comment")) == false {
				http.Error(w, "更新に失敗しました。入力文字数が最大値を超えている可能性があります。", http.StatusInternalServerError)
				RedirectTo(w, fmt.Sprintf("/user/%d/profile", userID))
				return nil
			}
			user := infrastructure.InfrastructureOBJ.UserAccesser.FindByID(userID)
			user.Name = r.FormValue("userName")
			user.Comment = r.FormValue("comment")
			if us.UpdateUser(user) == false {
				http.Error(w, "データベースに登録できませんでした。", http.StatusInternalServerError)
			}
			RedirectTo(w, fmt.Sprintf("/user/%d/profile", userID))
		}
	case "articles":
	}

	return nil
}

// SetUserToCookie ユーザ情報をcookieにセットします
func SetUserToCookie(w http.ResponseWriter, cookieName string, value *model.User) {
	authCookieValue := objx.New(map[string]interface{}{
		"user":   value,
		"userID": value.ID,
	}).MustBase64()
	http.SetCookie(w, &http.Cookie{
		Name:  cookieName,
		Value: authCookieValue,
		Path:  "/"}) //TODO: Pathが"/"なのはセキュリティ上よくないので、厳密に指定する。
}

func (c *UserController) specifyTemplate() *template.Template {
	return templateHelperOBJ.compiledTemplates[c.htmlFilename]
}
