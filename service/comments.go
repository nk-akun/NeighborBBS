package service

import (
	"github.com/nk-akun/NeighborBBS/model"
	"github.com/nk-akun/NeighborBBS/repository"
	"github.com/nk-akun/NeighborBBS/util"
)

type commentService struct {
}

// CommentService is the entrance as a convenient interface
var CommentService = newCommentService()

func newCommentService() *commentService {
	return new(commentService)
}

func (s *commentService) BuildComment(userID int64, articleID int64, parentID int64, content string) (*model.Comment, error) {
	comment := &model.Comment{
		UserID:     userID,
		ArticleID:  articleID,
		ParentID:   parentID,
		Content:    content,
		CreateTime: util.NowTimestamp(),
	}

	if err := repository.CommentRepository.Create(util.DB(), comment); err != nil {
		return nil, err
	}
	return comment, nil
}
