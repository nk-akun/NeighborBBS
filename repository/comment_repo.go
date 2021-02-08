package repository

import (
	"github.com/nk-akun/NeighborBBS/logs"
	"github.com/nk-akun/NeighborBBS/model"
	"gorm.io/gorm"
)

type commentRepository struct {
}

// CommentRepository is the entrance as a convenient interface
var CommentRepository = newCommentRepository()

func newCommentRepository() *commentRepository {
	return new(commentRepository)
}

func (r *commentRepository) Create(db *gorm.DB, comment *model.Comment) error {
	if err := db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (r *commentRepository) GetCommentsByCursorTime(db *gorm.DB, articleID int64, cursorTime int64) ([]model.Comment, error) {
	var comments []model.Comment
	err := db.Where("create_time < ?", cursorTime).Where("article_id = ?", articleID).Limit(30).Find(&comments).Error
	if err != nil {
		logs.Logger.Errorf("query db error:", err)
		return nil, err
	}
	return comments, err
}

func (r *commentRepository) GetCommentsByCommentID(db *gorm.DB, id int64) (*model.Comment, error) {
	return r.takeOne(db, "id = ?", id)
}

func (r *commentRepository) GetCommentsByArticleID(db *gorm.DB, articleID int64) ([]model.Comment, error) {
	return r.takeList(db, "article_id = ?", articleID)
}

func (r *commentRepository) takeOne(db *gorm.DB, column string, value interface{}) (*model.Comment, error) {
	comment := &model.Comment{}
	err := db.Where(column, value).Find(&comment).Error
	if err != nil {
		logs.Logger.Errorf("query db error:", err)
		return nil, err
	}
	return comment, nil
}

func (r *commentRepository) takeList(db *gorm.DB, column string, value interface{}) ([]model.Comment, error) {
	var comments []model.Comment
	err := db.Where(column, value).Find(&comments).Error
	if err != nil {
		logs.Logger.Errorf("query db error:", err)
		return nil, err
	}
	return comments, nil
}
