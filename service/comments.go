package service

import (
	"errors"

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

func (s *commentService) GetCommentList(articleID int64) (*model.CommentListResponse, error) {
	resp := new(model.CommentListResponse)
	comtList, err := repository.CommentRepository.GetCommentsByArticleID(util.DB(), articleID)
	if err != nil {
		return nil, errors.New("查询评论信息出错")
	}

	resp.ArticleID = articleID
	resp.TotalNum = len(comtList)
	resp.CommentList = make([]*model.CommentInfo, len(comtList))
	for i := range comtList {
		userInfo := repository.UserRepository.GetUserByUserID(util.DB(), comtList[i].UserID)
		if userInfo == nil {
			return nil, errors.New("查询作者信息出错")
		}
		resp.CommentList[i] = &model.CommentInfo{
			AuthorName: userInfo.Nickname,
			AuthorID:   userInfo.ID,
			AvatarURL:  userInfo.AvatarURL,
			Content:    comtList[i].Content,
			ParentID:   comtList[i].ParentID,
			LikeCount:  comtList[i].LikeCount,
			CreateTime: comtList[i].CreateTime,
		}
	}
	return resp, nil
}
