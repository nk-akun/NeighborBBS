package service

import (
	"errors"

	"github.com/nk-akun/NeighborBBS/model"
	"github.com/nk-akun/NeighborBBS/repository"
	"github.com/nk-akun/NeighborBBS/util"
	"gorm.io/gorm"
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
		return errors.New("操作失败")
	}
	if opHis.Status == 1 {
		return errors.New("已点赞，请勿重复操作")
	}

	err = util.DB().Transaction(func(tx *gorm.DB) error {
		var err error
		if opHis.ID != 0 {
			if err = repository.LCRepository.UpdateUserLikeOperation(util.DB(), userID, articleID, map[string]interface{}{"status": 1}); err != nil {
				return err
			}
		} else {
			err = repository.LCRepository.Create(util.DB(), &model.UserLikeArticle{
				UserID:     userID,
				ArticleID:  articleID,
				Status:     1,
				UpdateTime: util.NowTimestamp(),
			})
			if err != nil {
				return err
			}
		}
		err = util.DB().Exec("update t_article set like_count = like_count+1 where user_id = ? and article_id = ?", userID, articleID).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
