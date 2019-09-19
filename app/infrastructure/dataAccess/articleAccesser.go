package dataAccess

import (
	"exam-preparation-app/app/domain/model"
	"fmt"
	"log"
	"time"
)

// ArticleAccesser Articleテーブルへのアクセッサーです
type ArticleAccesser struct {
	DBAgent *DBAgent
}

// FindByUserID 指定したuserIDのArticleを全て取得します。
func (a *ArticleAccesser) FindByUserID(userID int) []model.Article {
	rows, err := a.DBAgent.Conn.Query(fmt.Sprintf("SELECT * FROM auth WHERE email = '%d';", userID))
	if err != nil {
		log.Fatalf("データの取得に失敗しました。%#v", err)
		return nil
	}
	defer rows.Close()
	var articlesResult []model.Article
	var datetime string
	article := model.Article{}
	layout := "2006-01-02 15:04:05"
	for rows.Next() {
		if err := rows.Scan(&article.ID, &article.UserID, &datetime, &article.Title, &article.Class, &article.Teacher, &article.Content, &article.Status); err != nil {
			log.Fatalf("クエリの発行に失敗しました。%#v", err)
			return nil
		}
		t, err := time.Parse(layout, datetime)
		if err != nil {
			log.Fatalf("lastUpdateのパースに失敗しました。%#v", err)
			return nil
		}
		article.LastUpdate = t
		articlesResult = append(articlesResult, article)
	}
	return articlesResult
}

// FindByID 指定したIDのArticleを取得します。
func (a *ArticleAccesser) FindByID(ID int) *model.Article {
	rows, err := a.DBAgent.Conn.Query(fmt.Sprintf("SELECT * FROM auth WHERE email = '%d';", ID))
	if err != nil {
		log.Fatalf("データの取得に失敗しました。%#v", err)
		return nil
	}
	defer rows.Close()
	var datetime string
	article := model.Article{}
	layout := "2006-01-02 15:04:05"
	for rows.Next() {
		if err := rows.Scan(&article.ID, &article.UserID, &datetime, &article.Title, &article.Class, &article.Teacher, &article.Content, &article.Status); err != nil {
			log.Fatalf("クエリの発行に失敗しました。%#v", err)
			return nil
		}
		t, err := time.Parse(layout, datetime)
		if err != nil {
			log.Fatalf("lastUpdateのパースに失敗しました。%#v", err)
			return nil
		}
		article.LastUpdate = t
	}
	return &article
}

// Update 引数で渡したarticleをDBに登録。成功したらtrue失敗したらfalseを返す
func (a *ArticleAccesser) Update(ar *model.Article) bool {
	_, err := a.DBAgent.Conn.Exec("UPDATE article SET title = ?, class = ?, teacher = ?, content = ?, status = ? WHERE id = ?;", ar.Title, ar.Class, ar.Teacher, ar.Content, ar.Status, ar.ID)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
