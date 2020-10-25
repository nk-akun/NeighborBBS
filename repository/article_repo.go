package repository

import (
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

// Create ...
func (r *articleRepository) Create(db *gorm.DB, article *model.Article) error {
	if err := db.Create(article).Error; err != nil {
		return err
	}
	return nil
}
