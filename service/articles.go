package service

import (
	"errors"

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
		return nil, err
	}
	return article, nil
}

func (s *articleService) GetArticleList(limit int, sortby string, order string) (*model.ArticleListResponse, error) {
	resp := &model.ArticleListResponse{}
	fields := []string{"id", "title", "create_time"}
	articles := repository.ArticleRepository.GetArticleFields(util.DB(), fields, limit, sortby, order)

	resp.TotalNum = len(articles)
	resp.ArticleList = make([]*model.ArticleBriefInfo, resp.TotalNum)
	for i := range articles {
		resp.ArticleList[i] = new(model.ArticleBriefInfo)
		resp.ArticleList[i].ArticleID = articles[i].ID
		resp.ArticleList[i].Title = articles[i].Title
		resp.ArticleList[i].CreateTime = articles[i].CreateTime
	}
	return resp, nil
}

func (s *articleService) GetArticleByID(id int64) (*model.ArticleResponse, error) {
	articleInfo := repository.ArticleRepository.GetArticleByID(util.DB(), id)
	if articleInfo == nil {
		return nil, errors.New("查询文章信息出错")
	}
	userInfo := repository.UserRepository.GetUserByUserID(util.DB(), articleInfo.UserID)
	if userInfo == nil {
		return nil, errors.New("查询作者信息出错")
	}

	resp := &model.ArticleResponse{
		Title:        articleInfo.Title,
		Author:       userInfo.Nickname,
		Content:      articleInfo.Content,
		CommentCount: articleInfo.CommentCount,
		LikeCount:    articleInfo.LikeCount,
		CreateTime:   articleInfo.CreateTime,
	}
	return resp, nil
}
