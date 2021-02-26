package service

import (
	"errors"
	"unicode/utf8"

	"github.com/nk-akun/NeighborBBS/model"
	"github.com/nk-akun/NeighborBBS/repository"
	"github.com/nk-akun/NeighborBBS/util"
	"gorm.io/gorm"
)

type articleService struct {
}

// ArticleService is the entrance as a convenient interface
var ArticleService = newArticleService()

func newArticleService() *articleService {
	return new(articleService)
}

func (s *articleService) BuildArticle(user *model.User, title string, content string) (*model.Article, error) {
	article := &model.Article{
		UserID:     user.ID,
		Title:      title,
		Content:    content,
		CreateTime: util.NowTimestamp(),
	}

	err := util.DB().Transaction(func(tx *gorm.DB) error {
		var err error

		err = repository.ArticleRepository.Create(util.DB(), article)
		if err != nil {
			return err
		}
		err = util.DB().Exec("update t_user set post_count = post_count+1 where id = ?", user.ID).Error
		return err
	})
	if err != nil {
		return nil, errors.New("数据库操作出错")
	}

	return article, nil
}

func (s *articleService) GetArticleList(currentUser *model.User, authorID int64, limit int, cursorTime int64, sortby string, order string) (*model.ArticleListResponse, error) {
	resp := &model.ArticleListResponse{}
	fields := []string{"id", "title", "create_time", "user_id", "view_count", "comment_count", "like_count", "content"}
	articles := repository.ArticleRepository.GetArticleFields(util.DB(), authorID, fields, cursorTime, limit, sortby, order)

	briefList, minCursorTime := buildArticleList(currentUser, articles)

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

func (s *articleService) GetArticleByID(currentUser *model.User, id int64) (*model.ArticleResponse, error) {
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
		Content:      util.MarkdownToHTML(articleInfo.Content),
		Liked:        LCService.JudgeArticleLiked(articleInfo, currentUser),
		CommentCount: articleInfo.CommentCount,
		LikeCount:    articleInfo.LikeCount,
		CreateTime:   articleInfo.CreateTime,
	}
	return resp, nil
}

func buildArticleList(currentUser *model.User, articles []model.Article) ([]*model.ArticleBriefInfo, int64) {
	var minCursorTime int64 = model.MAXCursorTime
	briefList := make([]*model.ArticleBriefInfo, len(articles))
	for i := range articles {
		minCursorTime = util.MinInt64(minCursorTime, articles[i].CreateTime)
		mkSummary := util.MarkdownToHTML(util.SubString(articles[i].Content, 0, util.MinInt(128, utf8.RuneCountInString(articles[i].Content))))
		briefList[i] = new(model.ArticleBriefInfo)
		briefList[i].ArticleID = articles[i].ID
		briefList[i].Title = articles[i].Title
		briefList[i].Summary = util.GetHTMLText(mkSummary)
		briefList[i].CommentCount = articles[i].CommentCount
		briefList[i].LikeCount = articles[i].LikeCount
		briefList[i].ViewCount = articles[i].ViewCount
		briefList[i].CreateTime = articles[i].CreateTime
		briefList[i].Liked = LCService.JudgeArticleLiked(&articles[i], currentUser)
		user, _ := repository.UserRepository.GetUserByUserID(util.DB(), articles[i].UserID)
		briefList[i].User = BuildUserBriefInfo(user)
	}

	return briefList, minCursorTime
}
