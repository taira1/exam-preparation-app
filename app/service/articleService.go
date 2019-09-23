package service

import (
	"database/sql"
	"exam-preparation-app/app/domain/model"
	"exam-preparation-app/app/infrastructure"
	"exam-preparation-app/app/infrastructure/dataAccess"
)

// ArticleService 記事サービスです
type ArticleService struct {
	validater *Validater
}

// ArticleStatusCodes 許容される記事のステータスコードです
var ArticleStatusCodes = []string{"public", "pending"}

// ValidateStatus Statusコードの妥当性を検証します。
func (s *ArticleService) ValidateStatus(str string) bool {
	return s.validater.ValidateMatchesFixedValues(str, ArticleStatusCodes)
}

// Update 引数で指定した記事を更新します。
func (s *ArticleService) Update(article *model.Article) bool {
	return infrastructure.InfrastructureOBJ.ArticleAccesser.Update(article)
}

// DeleteByID idで指定した記事を削除します。
func (s *ArticleService) DeleteByID(id int) bool {
	return infrastructure.InfrastructureOBJ.ArticleAccesser.DeleteByID(id)
}

// NewArticleService コンストラクタです
func NewArticleService() *ArticleService {
	return &ArticleService{validater: NewValidater()}
}

// RegisterArticle 引数で指定したarticleをDBに保存します。
func (s *ArticleService) RegisterArticle(article *model.Article) {
	dataAccess.Transact(infrastructure.InfrastructureOBJ.DBAgent.Conn, func(tx *sql.Tx) error {
		return infrastructure.InfrastructureOBJ.ArticleAccesser.Insert(article)
	})

}
