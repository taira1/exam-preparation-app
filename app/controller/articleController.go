package controller

import (
	"exam-preparation-app/app/infrastructure"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// ArticleController グローバル記事コントローラです。
type ArticleController struct {
	htmlFilename string
}

func (c *ArticleController) process(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	segs := strings.Split(r.URL.Path, "/")
	var articleID = 0
	if len(segs) == 3 {
		articleID, _ = strconv.Atoi(segs[2])
	}
	switch articleID {
	case 0:
		// articleをキーUserを値にもつMapを取得する
		// articles := map[model.Article]*model.User{}
		// for _, a := range infrastructure.InfrastructureOBJ.ArticleAccesser.FindByStatusIsPublic() {
		// 	articles[a] = infrastructure.InfrastructureOBJ.UserAccesser.FindByID(a.UserID)
		// }
		c.htmlFilename = "new_article_list.html"
		return map[string]interface{}{
			"Articles": infrastructure.InfrastructureOBJ.ArticleAccesser.FindByStatusIsPublic(),
		}
	default:
		c.htmlFilename = "article.html"
		article := infrastructure.InfrastructureOBJ.ArticleAccesser.FindByID(articleID)
		return map[string]interface{}{
			"Article": article,
			"Poster":  infrastructure.InfrastructureOBJ.UserAccesser.FindByID(article.UserID),
		}
	}
}

func (c *ArticleController) specifyTemplate() *template.Template {
	return templateHelperOBJ.compiledTemplates[c.htmlFilename]
}
