package service

import (
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
	articleResp := &model.ArticleListResponse{}
	fields := []string{"id", "title", "create_time"}
	articles := repository.ArticleRepository.GetArticleFields(util.DB(), fields, limit, sortby, order)

	articleResp.TotalNum = len(articles)
	articleResp.ArticleList = make([]*model.ArticleBriefInfo, articleResp.TotalNum)
	for i := range articles {
		articleResp.ArticleList[i] = new(model.ArticleBriefInfo)
		articleResp.ArticleList[i].ArticleID = articles[i].ID
		articleResp.ArticleList[i].Title = articles[i].Title
		articleResp.ArticleList[i].CreateTime = articles[i].CreateTime
	}
	return articleResp, nil
}
