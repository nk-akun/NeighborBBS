package repository

import (
	"github.com/nk-akun/NeighborBBS/logs"
	"github.com/nk-akun/NeighborBBS/model"
	"gorm.io/gorm"
)

type userTokenRepository struct {
}

// UserTokenRepository is the entrance as a convenient interface
var UserTokenRepository = newUserTokenRepository()

func newUserTokenRepository() *userTokenRepository {
	return new(userTokenRepository)
}

func (r *userTokenRepository) Create(db *gorm.DB, userToken *model.UserToken) error {
	return db.Create(userToken).Error
}

func (r *userTokenRepository) GetUserIDByToken(db *gorm.DB, token string) (*model.UserToken, error) {
	return r.take(db, "token = ?", token)
}

func (r *userTokenRepository) take(db *gorm.DB, column string, value interface{}) (*model.UserToken, error) {
	result := &model.UserToken{}
	// err := db.Where(column, value).Take(result).Error
	err := db.Where(column, value).Find(&result).Error
	if err != nil {
		logs.Logger.Errorf("query db error:", err)
		return nil, err
	}
	return result, nil
}
