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

func (s *articleService) GetArticleList(limit int, sortby string, order string) {
	// articleResp := &model.ArticleListResponse{}

}
