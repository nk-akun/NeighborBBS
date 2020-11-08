package repository

import (
	"github.com/nk-akun/NeighborBBS/logs"
	"github.com/nk-akun/NeighborBBS/model"
	"gorm.io/gorm"
)

type lcRepository struct {
}

// LCRepository is the entrance as a convenient interface
var LCRepository = newLCRepository()

func newLCRepository() *lcRepository {
	return new(lcRepository)
}

func (r *lcRepository) Create(db *gorm.DB, op *model.UserLikeArticle) error {
	return db.Create(op).Error
}

func (r *lcRepository) GetUserLikeOperation(db *gorm.DB, userID int64, articleID int64) (*model.UserLikeArticle, error) {
	return r.takeOne(db, map[string]interface{}{
		"user_id":    userID,
		"article_id": articleID,
	})
}

func (r *lcRepository) takeOne(db *gorm.DB, params map[string]interface{}) (*model.UserLikeArticle, error) {
	opHis := &model.UserLikeArticle{}
	err := db.Where(params).Limit(1).Find(&opHis).Error
	if err != nil {
		logs.Logger.Error("query db error:", err)
		return nil, err
	}
	return opHis, nil
}

func (r *lcRepository) UpdateUserLikeOperation(db *gorm.DB, userID int64, articleID int64, params map[string]interface{}) error {
	return db.Model(&model.UserLikeArticle{}).Where("user_id = ? and article_id = ?", userID, articleID).Updates(params).Error
}
