package service

import (
	"exam-preparation-app/app/domain/model"
	"exam-preparation-app/app/infrastructure"
)

// ArticleService 記事サービスです
type ArticleService struct {
	validater *Validater
}

// ValidateStatus Statusコードの妥当性を検証します。
func (s *ArticleService) ValidateStatus(str string) bool {
	values := []string{"public", "pending"}
	return s.validater.ValidateMatchesFixedValues(str, values)
}

// Update 引数で指定した記事を更新します。
func (s *ArticleService) Update(article *model.Article) bool {
	return infrastructure.InfrastructureOBJ.ArticleAccesser.Update(article)
}

// NewArticleService コンストラクタです
func NewArticleService() *ArticleService {
	return &ArticleService{validater: NewValidater()}
}
