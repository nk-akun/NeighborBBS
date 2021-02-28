package service

import (
	"errors"

	"github.com/nk-akun/NeighborBBS/model"
	"github.com/nk-akun/NeighborBBS/repository"
	"github.com/nk-akun/NeighborBBS/util"
	"gorm.io/gorm"
)

const (
	LikeArticle     = 1
	FavoriteArticle = 1
)

type lcService struct {
}

// LCService is the entrance as a convenient interface
var LCService = newLCService()

func newLCService() *lcService {
	return new(lcService)
}

func (s *lcService) PostLikeArticle(userID int64, articleID int64) error {
	opHis, err := repository.LCRepository.GetUserLikeOperation(util.DB(), userID, articleID)
	if err != nil {
		return errors.New("数据库查询失败")
	}

	//即已经点赞成功
	if opHis.Status == 1 {
		return nil
	}

	err = util.DB().Transaction(func(tx *gorm.DB) error {
		var err error
		if opHis.ID != 0 {
			err = repository.LCRepository.UpdateUserLikeOperation(util.DB(), userID, articleID, map[string]interface{}{"status": 1})
		} else {
			err = repository.LCRepository.CreateLike(util.DB(), &model.UserLikeArticle{
				UserID:     userID,
				ArticleID:  articleID,
				Status:     1,
				UpdateTime: util.NowTimestamp(),
			})
		}
		if err != nil {
			return err
		}
		err = util.DB().Exec("update t_article set like_count = like_count+1 where id = ?", articleID).Error
		return err
	})
	if err != nil {
		return errors.New("数据库操作出错")
	}
	return nil
}

// JudgeArticleLiked judge user has liked the article or not.
func (s *lcService) JudgeArticleLiked(article *model.Article, user *model.User) bool {
	if user == nil {
		return false
	}
	lcStatus, _ := repository.LCRepository.GetUserLikeOperation(util.DB(), user.ID, article.ID)
	return lcStatus.Status == LikeArticle
}

// JudgeArticleFavorited judge user has liked the article or not.
func (s *lcService) JudgeArticleFavorited(article *model.Article, user *model.User) bool {
	if user == nil {
		return false
	}
	lcStatus, _ := repository.LCRepository.GetUserFavoriteOperation(util.DB(), user.ID, article.ID)
	return lcStatus.Status == FavoriteArticle
}

func (s *lcService) PostDelLikeArticle(userID int64, articleID int64) error {
	opHis, err := repository.LCRepository.GetUserLikeOperation(util.DB(), userID, articleID)
	if err != nil || opHis.ID == 0 {
		return errors.New("数据库查询失败")
	}

	//即已经取消点赞
	if opHis.Status == 0 {
		return nil
	}

	err = util.DB().Transaction(func(tx *gorm.DB) error {
		var err error

		err = repository.LCRepository.UpdateUserLikeOperation(util.DB(), userID, articleID, map[string]interface{}{"status": 0})
		if err != nil {
			return err
		}
		err = util.DB().Exec("update t_article set like_count = like_count-1 where id = ?", articleID).Error
		return err
	})
	if err != nil {
		return errors.New("数据库操作出错")
	}
	return nil
}

func (s *lcService) PostFavoriteArticle(userID int64, articleID int64) error {
	opHis, err := repository.LCRepository.GetUserFavoriteOperation(util.DB(), userID, articleID)
	if err != nil {
		return errors.New("数据库查询失败")
	}

	//即已经收藏成功
	if opHis.Status == 1 {
		return nil
	}

	err = util.DB().Transaction(func(tx *gorm.DB) error {
		var err error
		if opHis.ID != 0 {
			err = repository.LCRepository.UpdateUserFavoriteOperation(util.DB(), userID, articleID, map[string]interface{}{"status": 1})
			util.DB().Exec("update t_user_favorite_article set update_time = ? where id = ?", util.NowTimestamp(), opHis.ID)
		} else {
			err = repository.LCRepository.CreateFavorite(util.DB(), &model.UserFavoriteArticle{
				UserID:     userID,
				ArticleID:  articleID,
				Status:     1,
				CreateTime: util.NowTimestamp(),
				UpdateTime: util.NowTimestamp(), //每次收藏或取消收藏都要更新此时间
			})
		}
		if err != nil {
			return err
		}
		err = util.DB().Exec("update t_user set favourite_article_count = favourite_article_count+1 where id = ?", userID).Error
		return err
	})
	if err != nil {
		return errors.New("数据库操作出错")
	}
	return nil
}

func (s *lcService) PostDelFavoriteArticle(userID int64, articleID int64) error {
	opHis, err := repository.LCRepository.GetUserFavoriteOperation(util.DB(), userID, articleID)
	if err != nil || opHis.ID == 0 {
		return errors.New("数据库查询失败")
	}

	//即已经取消收藏
	if opHis.Status == 0 {
		return nil
	}

	err = util.DB().Transaction(func(tx *gorm.DB) error {
		var err error

		err = repository.LCRepository.UpdateUserFavoriteOperation(util.DB(), userID, articleID, map[string]interface{}{"status": 0})
		util.DB().Exec("update t_user_favorite_article set update_time = ? where id = ?", util.NowTimestamp(), opHis.ID)
		if err != nil {
			return err
		}
		err = util.DB().Exec("update t_user set favourite_article_count = favourite_article_count-1 where id = ?", userID).Error
		return err
	})
	if err != nil {
		return errors.New("数据库操作出错")
	}
	return nil
}

func (s *lcService) GetFavoriteArticles(user *model.User, limit int, cursorTime int64, sortby string, order string) (*model.FavoriteResponse, error) {
	records := repository.LCRepository.GetFavoriteRecords(util.DB(), user.ID, cursorTime, limit, sortby, order)

	articles := make([]model.Article, 0, len(records))
	for _, record := range records {
		article, _ := repository.ArticleRepository.GetArticleByID(util.DB(), record.ArticleID)
		articles = append(articles, *article)
	}

	briefList, _ := ArticleService.BuildArticleList(user, articles)
	minCursorTime := cursorTime
	for i := range records {
		minCursorTime = util.MinInt64(minCursorTime, records[i].UpdateTime)
	}
	resp := &model.FavoriteResponse{}
	resp.TotalNum = len(briefList)
	resp.Cursor = minCursorTime
	resp.FavoriteList = briefList
	return resp, nil
}
