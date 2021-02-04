package repository

import (
	"fmt"

	"github.com/nk-akun/NeighborBBS/logs"
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
	return db.Create(article).Error
}

func (r *articleRepository) GetArticleFields(db *gorm.DB, fields []string, cursorTime int64, limit int, sortby string, order string) []model.Article {
	var articles []model.Article
	db.Where("create_time < ?", cursorTime).Select(fields).Order(fmt.Sprintf("%s %s", sortby, order)).Limit(limit).Find(&articles)
	return articles
}

func (r *articleRepository) GetArticleByID(db *gorm.DB, id int64) (*model.Article, error) {
	return r.take(db, "id = ?", id)
}

func (r *articleRepository) take(db *gorm.DB, column string, value interface{}) (*model.Article, error) {
	result := new(model.Article)
	if err := db.Where(column, value).Find(&result).Error; err != nil {
		logs.Logger.Errorf("query db error:", err)
		return nil, err
	}
	return result, nil
}
