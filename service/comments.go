package service

import (
	"errors"

	"github.com/nk-akun/NeighborBBS/logs"
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
		logs.Logger.Error("db error:", err)
		return nil, errors.New("数据库操作出错")
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
	buildCommentList(comtList)
}

func buildCommentList(comtList []model.Comment) []*model.CommentInfo {
	// TODO: 对comtList排序，然后构造ParentComment
	// https://itimetraveler.github.io/2016/09/07/%E3%80%90Go%E8%AF%AD%E8%A8%80%E3%80%91%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B%E6%8E%92%E5%BA%8F%E5%92%8C%20slice%20%E6%8E%92%E5%BA%8F/
	detailedCommentList := make([]*model.CommentInfo, len(comtList))
	for i := range comtList {
		userInfo, err := repository.UserRepository.GetUserByUserID(util.DB(), comtList[i].UserID)
		if err != nil {
			logs.Logger.Errorf("查询作者信息出错")
		}
		detailedCommentList[i] = &model.CommentInfo{
			AuthorNickName: userInfo.Nickname,
			AuthorUserName: userInfo.Username,
			AuthorID:       userInfo.ID,
			AvatarURL:      userInfo.AvatarURL,
			Content:        comtList[i].Content,
			ParentComment: 
			LikeCount:      comtList[i].LikeCount,
			CreateTime:     comtList[i].CreateTime,
		}
	}
}
