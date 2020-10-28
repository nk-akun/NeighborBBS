package repository

import (
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
