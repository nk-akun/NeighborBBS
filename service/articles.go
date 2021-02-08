package service

import (
	"errors"

	"github.com/nk-akun/NeighborBBS/logs"
	"github.com/nk-akun/NeighborBBS/model"
	"github.com/nk-akun/NeighborBBS/repository"
	"github.com/nk-akun/NeighborBBS/util"
)

type articleService struct {
}

// ArticleService is the entrance as a convenient interface
var ArticleService = newArticleService()

func newArticleService() *articleService {
	return new(articleService)
}

func (s *articleService) BuildArticle(userID int64, title string, content string) (*model.Article, error) {
	article := &model.Article{
		UserID:     userID,
		Title:      title,
		Content:    content,
		CreateTime: util.NowTimestamp(),
	}

	if err := repository.ArticleRepository.Create(util.DB(), article); err != nil {
		logs.Logger.Error("db error:", err)
		return nil, errors.New("数据库操作出错")
	}
	return article, nil
}

func (s *articleService) GetArticleList(limit int, cursorTime int64, sortby string, order string) (*model.ArticleListResponse, error) {
	resp := &model.ArticleListResponse{}
	fields := []string{"id", "title", "create_time", "user_id", "view_count", "comment_count", "like_count"}
	articles := repository.ArticleRepository.GetArticleFields(util.DB(), fields, cursorTime, limit, sortby, order)

	briefList, minCursorTime := buildArticleList(articles)

	resp.Cursor = minCursorTime
	for i := range briefList {
		if briefList[i].CreateTime < resp.Cursor {
			resp.Cursor = briefList[i].CreateTime
		}
	}
	resp.ArticleList = briefList
	resp.TotalNum = len(briefList)
	return resp, nil
}

func (s *articleService) GetArticleByID(id int64) (*model.ArticleResponse, error) {
	articleInfo, err := repository.ArticleRepository.GetArticleByID(util.DB(), id)
	if err != nil {
		return nil, errors.New("查询文章信息出错")
	}
	userInfo, err := repository.UserRepository.GetUserByUserID(util.DB(), articleInfo.UserID)
	if err != nil {
		return nil, errors.New("查询作者信息出错")
	}

	resp := &model.ArticleResponse{
		ArticleID:    articleInfo.ID,
		Title:        articleInfo.Title,
		User:         BuildUserBriefInfo(userInfo),
		Content:      util.ToHTML(articleInfo.Content),
		CommentCount: articleInfo.CommentCount,
		LikeCount:    articleInfo.LikeCount,
		CreateTime:   articleInfo.CreateTime,
	}
	return resp, nil
}

func buildArticleList(articles []model.Article) ([]*model.ArticleBriefInfo, int64) {
	var minCursorTime int64 = model.MAXCursorTime
	briefList := make([]*model.ArticleBriefInfo, len(articles))
	for i := range articles {
		minCursorTime = util.MinInt64(minCursorTime, articles[i].CreateTime)
		briefList[i] = new(model.ArticleBriefInfo)
		briefList[i].ArticleID = articles[i].ID
		briefList[i].Title = articles[i].Title
		briefList[i].CommentCount = articles[i].CommentCount
		briefList[i].LikeCount = articles[i].LikeCount
		briefList[i].ViewCount = articles[i].ViewCount
		briefList[i].CreateTime = articles[i].CreateTime
		user, _ := repository.UserRepository.GetUserByUserID(util.DB(), articles[i].UserID)
		briefList[i].Liked = LCService.JudgeArticleLiked(&articles[i], user)
		briefList[i].User = BuildUserBriefInfo(user)
	}

	return briefList, minCursorTime
}
