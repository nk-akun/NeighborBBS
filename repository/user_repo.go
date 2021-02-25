package repository

import (
	"github.com/nk-akun/NeighborBBS/logs"
	"github.com/nk-akun/NeighborBBS/model"
	"gorm.io/gorm"
)

type userRepository struct {
}

// UserRepository is the entrance as a convenient interface
var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return new(userRepository)
}

// Create ...
func (r *userRepository) Create(db *gorm.DB, user *model.User) error {
	return db.Create(user).Error
}

func (r *userRepository) GetUserByEmail(db *gorm.DB, email string) (*model.User, error) {
	return r.take(db, "email = ?", email)
}

func (r *userRepository) GetUserByUsername(db *gorm.DB, username string) (*model.User, error) {
	return r.take(db, "username = ?", username)
}

func (r *userRepository) GetUserByUserID(db *gorm.DB, userID int64) (*model.User, error) {
	return r.take(db, "id = ?", userID)
}

func (r *userRepository) UpdateOne(db *gorm.DB, userID int64, column string, value interface{}) error {
	return db.Model(model.User{}).Where("id = ?", userID).Update(column, value).Error
}

func (r *userRepository) UpdateMulti(db *gorm.DB, userID int64, kv map[string]interface{}) error {
	return db.Model(model.User{}).Where("id = ?", userID).Updates(kv).Error
}

func (r *userRepository) take(db *gorm.DB, column string, value interface{}) (*model.User, error) {
	result := &model.User{}
	// err := db.Where(column, value).Take(result).Error
	err := db.Where(column, value).Find(&result).Error
	if err != nil {
		logs.Logger.Errorf("query db error:", err)
		return nil, err
	}
	return result, nil
}
