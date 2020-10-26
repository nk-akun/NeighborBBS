package repository

import (
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
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUserByEmail(db *gorm.DB, email string) *model.User {
	return r.take(db, "email = ?", email)
}

func (r *userRepository) GetUserByUsername(db *gorm.DB, username string) *model.User {
	return r.take(db, "username = ?", username)
}

func (r *userRepository) take(db *gorm.DB, column string, value interface{}) *model.User {
	result := &model.User{}
	// err := db.Where(column, value).Take(result).Error
	err := db.Where(column, value).Find(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return result
}
