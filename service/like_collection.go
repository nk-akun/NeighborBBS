package service

import (
	"errors"

	"github.com/nk-akun/NeighborBBS/model"
	"github.com/nk-akun/NeighborBBS/repository"
	"github.com/nk-akun/NeighborBBS/util"
	"gorm.io/gorm"
)

const (
	LikeArticle = 1
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
			err = repository.LCRepository.Create(util.DB(), &model.UserLikeArticle{
				UserID:     userID,
				ArticleID:  articleID,
				Status:     1,
				UpdateTime: util.NowTimestamp(),
			})
		}
		if err != nil {
			return err
		}
		err = util.DB().Exec("update t_article set like_count = like_count+1 where id = ? and user_id = ?", articleID, userID).Error
		return err
	})
	if err != nil {
		return errors.New("数据库操作出错")
	}
	return nil
}

// JudgeArticleLiked judge user has liked the article or not.
func (s *lcService) JudgeArticleLiked(article *model.Article, user *model.User) bool {
	lcStatus, _ := repository.LCRepository.GetUserLikeOperation(util.DB(), user.ID, article.ID)
	return lcStatus.Status == LikeArticle
}

func (s *lcService) PostDelLikeArticle(userID int64, articleID int64) error {
	opHis, err := repository.LCRepository.GetUserLikeOperation(util.DB(), userID, articleID)
	if err != nil || opHis.ID == 0 {
		return errors.New("数据库查询失败")
	}

	//即已经取消成功
	if opHis.Status == 0 {
		return nil
	}

	err = util.DB().Transaction(func(tx *gorm.DB) error {
		var err error

		err = repository.LCRepository.UpdateUserLikeOperation(util.DB(), userID, articleID, map[string]interface{}{"status": 0})
		if err != nil {
			return err
		}
		err = util.DB().Exec("update t_article set like_count = like_count-1 where id = ? and user_id = ?", articleID, userID).Error
		return err
	})
	if err != nil {
		return errors.New("数据库操作出错")
	}
	return nil
}
