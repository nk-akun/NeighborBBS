package repository

import (
	"fmt"

	"github.com/nk-akun/NeighborBBS/model"
	"gorm.io/gorm"
)

type articleRepository struct {
}

// ArticleRepository is the entrance as a convenient interface
var ArticleRepository = newArticleRepository()

func newArticleRepository() *articleRepository {
	return new(articleRepository)
}

func (r *articleRepository) Create(db *gorm.DB, article *model.Article) error {
	if err := db.Create(article).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) GetArticleFields(db *gorm.DB, fields []string, limit int, sortby string, order string) []model.Article {
	var articles []model.Article
	db.Select(fields).Order(fmt.Sprintf("%s %s", sortby, order)).Limit(limit).Find(&articles)
	return articles
}
