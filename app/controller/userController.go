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
	var method2 string
	if len(segs) == 5 {
		method = segs[4]
	}
	if len(segs) == 6 {
		method = segs[4]
		method2 = segs[5]
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
			c.htmlFilename = "user_edit.html"
			return map[string]interface{}{
				"User": infrastructure.InfrastructureOBJ.UserAccesser.FindByID(userID),
			}
		case "post":
			us := service.NewUserService()
			if us.ValidateUpdateUser(r.FormValue("userName"), r.FormValue("comment")) == false {
				http.Error(w, "更新に失敗しました。入力文字数が最大値を超えている可能性があります。", http.StatusInternalServerError)
				http.Redirect(w, r, fmt.Sprintf("/user/%d/profile/edit", userID), http.StatusTemporaryRedirect)
				return nil
			}
			user := infrastructure.InfrastructureOBJ.UserAccesser.FindByID(userID)
			user.Name = r.FormValue("userName")
			user.Comment = r.FormValue("comment")
			if us.UpdateUser(user) == false {
				http.Error(w, "データベースに登録できませんでした。", http.StatusInternalServerError)
			}
			http.Redirect(w, r, fmt.Sprintf("/user/%d/profile", userID), http.StatusTemporaryRedirect)
		}
	case "articles":
		switch method {
		case "": //記事一覧表示
			c.htmlFilename = "user_article_list.html"
			return map[string]interface{}{
				"Articles": infrastructure.InfrastructureOBJ.ArticleAccesser.FindByID(userID),
				"User":     infrastructure.InfrastructureOBJ.UserAccesser.FindByID(userID),
			}
		case "new": //記事の作成
			switch method2 {
			case "": //記事新規作成
				c.htmlFilename = "user_article_edit.html"
				article := &model.Article{
					Title:   r.FormValue("title"),
					Class:   r.FormValue("class"),
					Teacher: r.FormValue("teacher"),
					Content: r.FormValue("content"),
					Status:  r.FormValue("status"),
				}
				return map[string]interface{}{
					"User":    infrastructure.InfrastructureOBJ.UserAccesser.FindByID(userID),
					"Article": article,
				}
			case "post": //記事新規作成post
				as := service.ArticleService{}
				if as.ValidateStatus(r.FormValue("status")) == false {
					http.Error(w, "不適切な記事ステータスです", http.StatusInternalServerError)
					http.Redirect(w, r, fmt.Sprintf("/user/%d/article/new", userID), 301)
				}
			case "complete": //記事新規作成コンプリート
				c.htmlFilename = "article_complete.html"
				return nil
			}

		default: //記事の編集
			articleID, _ := strconv.Atoi(segs[4])
			switch method2 {
			case "edit": //記事編集
				c.htmlFilename = "user_article_edit.html"
				return map[string]interface{}{
					"Article": infrastructure.InfrastructureOBJ.ArticleAccesser.FindByID(articleID),
					"User":    infrastructure.InfrastructureOBJ.UserAccesser.FindByID(userID),
				}
			case "post": //記事編集のpost
				//TODO: ロジックの実装
				as := service.ArticleService{}
				if as.ValidateStatus(r.FormValue("status")) == false {
					http.Error(w, "不適切な記事ステータスです", http.StatusInternalServerError)
					http.Redirect(w, r, fmt.Sprintf("/user/%d/article/new", userID), 301)
				}
				article := &model.Article{
					ID:      articleID,
					UserID:  userID,
					Title:   r.FormValue("title"),
					Class:   r.FormValue("class"),
					Teacher: r.FormValue("teacher"),
					Content: r.FormValue("content"),
					Status:  r.FormValue("status"),
				}
				if as.Update(article) == false {
					http.Error(w, "記事の更新に失敗しました。", http.StatusInternalServerError)
					http.Redirect(w, r, fmt.Sprintf("/user/%d/articles/%d/edit", userID, articleID), http.StatusTemporaryRedirect)
					return nil
				}
				http.Redirect(w, r, fmt.Sprintf("/user/%d/article/%d/complete", userID, articleID), http.StatusTemporaryRedirect)
			case "complete": //記事新規編集コンプリート
				c.htmlFilename = "article_complete.html"
				return nil
			}
		}
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
