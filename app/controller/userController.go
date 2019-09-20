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
				"Err":  getErrMessagesFromCookie(r),
			}
		case "post":
			us := service.NewUserService()
			if us.ValidateUpdateUser(r.FormValue("userName"), r.FormValue("comment")) == false {
				setErrMessageToCookie(w, r, "更新に失敗しました。入力文字数が最大値を超えている可能性があります。")
				http.Redirect(w, r, fmt.Sprintf("/user/%d/profile/edit", userID), http.StatusTemporaryRedirect)
				return nil
			}
			user := infrastructure.InfrastructureOBJ.UserAccesser.FindByID(userID)
			user.Name = r.FormValue("userName")
			user.Comment = r.FormValue("comment")
			if us.UpdateUser(user) == false {
				setErrMessageToCookie(w, r, "データベースに登録できませんでした")
			}
			http.Redirect(w, r, fmt.Sprintf("/user/%d/profile", userID), http.StatusTemporaryRedirect)
		}
	case "article":
		switch method {
		case "": //記事一覧表示
			c.htmlFilename = "article_list.html"
			return map[string]interface{}{
				"Articles": infrastructure.InfrastructureOBJ.ArticleAccesser.FindByID(userID),
				"User":     infrastructure.InfrastructureOBJ.UserAccesser.FindByID(userID),
			}
		case "new": //記事の作成
			switch method2 {
			case "": //記事新規作成
				c.htmlFilename = "article_edit.html"
				data := map[string]interface{}{
					"User":          infrastructure.InfrastructureOBJ.UserAccesser.FindByID(userID),
					"Article":       getArticleFromCookie(r),
					"ArticleStatus": service.ArticleStatusCodes,
					"Next":          map[string]string{"URL": fmt.Sprintf("/user/%d/article/new/post", userID), "Status": "新規作成"},
					"Err":           getErrMessagesFromCookie(r),
				}
				deleteCookieByName(w, r, "article")
				deleteCookieByName(w, r, "err")
				return data
			case "post": //記事新規作成post
				as := service.ArticleService{}
				article := getArticleFromForm(r)
				article.UserID = userID
				if as.ValidateStatus(r.FormValue("status")) == false {
					setArticleToCookie(w, r, article)
					setErrMessageToCookie(w, r, "無効な記事ステータスです")
					http.Redirect(w, r, fmt.Sprintf("/user/%d/article/new", userID), 301)
					return nil
				}
				as.RegisterArticle(article)
				deleteCookieByName(w, r, "article")
				deleteCookieByName(w, r, "err")
				http.Redirect(w, r, fmt.Sprintf("/user/%d/article/new/complete", userID), 301)
			case "complete": //記事新規作成コンプリート
				c.htmlFilename = "article_complete.html"
				return nil
			}

		default: //記事の編集
			articleID, _ := strconv.Atoi(segs[4])
			switch method2 {
			case "edit": //記事編集
				c.htmlFilename = "article_edit.html"
				return map[string]interface{}{
					"Article": infrastructure.InfrastructureOBJ.ArticleAccesser.FindByID(articleID),
					"User":    infrastructure.InfrastructureOBJ.UserAccesser.FindByID(userID),
					"Next":    map[string]string{"URL": fmt.Sprintf("/user/%d/article/%d/post", userID, articleID), "Status": "記事編集"},
					"Err":     getErrMessagesFromCookie(r),
				}
			case "post": //記事編集のpost
				//TODO: ロジックの実装
				as := service.ArticleService{}
				if as.ValidateStatus(r.FormValue("status")) == false {
					setErrMessageToCookie(w, r, "無効な記事ステータスです")
					http.Redirect(w, r, fmt.Sprintf("/user/%d/article/new", userID), 301)
					return nil
				}
				article := getArticleFromForm(r)
				article.ID = articleID
				article.UserID = userID
				if as.Update(article) == false {
					setErrMessageToCookie(w, r, "BDの更新に失敗しました")
					http.Redirect(w, r, fmt.Sprintf("/user/%d/articles/%d/edit", userID, articleID), http.StatusTemporaryRedirect)
					return nil
				}
				deleteCookieByName(w, r, "err")
				http.Redirect(w, r, fmt.Sprintf("/user/%d/article/%d/complete", userID, articleID), http.StatusTemporaryRedirect)
			case "complete": //記事新規編集コンプリート
				c.htmlFilename = "article_complete.html"
				return nil
			}
		}
	}
	return nil
}

// setUserToCookie ユーザ情報をcookieにセットします
func setUserToCookie(w http.ResponseWriter, cookieName string, value *model.User) {
	userCookieValue := objx.New(map[string]interface{}{
		"user":   value,
		"userID": value.ID,
	}).MustBase64()
	http.SetCookie(w, &http.Cookie{
		Name:  cookieName,
		Value: userCookieValue,
		Path:  "/"}) //TODO: Pathが"/"なのはセキュリティ上よくないので、厳密に指定する。
}

// setArticleToCookie 記事情報をCookieにセットします
func setArticleToCookie(w http.ResponseWriter, r *http.Request, value *model.Article) {
	ArticleCookieValue := objx.New(map[string]interface{}{
		"Title":   value.Title,
		"Class":   value.Class,
		"Teacher": value.Teacher,
		"Content": value.Content,
		"Status":  value.Status,
	}).MustBase64()
	if cookieValue, err := r.Cookie("article"); err == nil {
		cookieValue.Value = ArticleCookieValue
		http.SetCookie(w, cookieValue)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "article",
		Value: ArticleCookieValue,
		Path:  "/"}) //TODO: Pathが"/"なのはセキュリティ上よくないので、厳密に指定する。
}

// getArticleFromCookie Cookieからarticleを取得します。
func getArticleFromCookie(r *http.Request) *model.Article {
	article := &model.Article{}
	if cookieValue, err := r.Cookie("article"); err == nil {
		if cookieValue.Value == "" {
			return article
		}
		if t, ok := objx.MustFromBase64(cookieValue.Value)["Title"].(string); ok {
			article.Title = t
		}
		if c, ok := objx.MustFromBase64(cookieValue.Value)["Class"].(string); ok {
			article.Class = c
		}
		if t, ok := objx.MustFromBase64(cookieValue.Value)["Teacher"].(string); ok {
			article.Teacher = t
		}
		if c, ok := objx.MustFromBase64(cookieValue.Value)["Content"].(string); ok {
			article.Content = c
		}
		if s, ok := objx.MustFromBase64(cookieValue.Value)["Status"].(string); ok {
			article.Status = s
		}
		return article
	}
	return nil
}

//getArticleFromForm
func getArticleFromForm(r *http.Request) *model.Article {
	return &model.Article{
		Title:   r.FormValue("title"),
		Class:   r.FormValue("class"),
		Teacher: r.FormValue("teacher"),
		Content: r.FormValue("content"),
		Status:  r.FormValue("status"),
	}
}

func (c *UserController) specifyTemplate() *template.Template {
	return templateHelperOBJ.compiledTemplates[c.htmlFilename]
}
