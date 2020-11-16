package guest

import (
	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/model"
	"github.com/nk-akun/NeighborBBS/service"
)

// PostLikeArticle post like
func PostLikeArticle(c *gin.Context) {
	req := getReqFromContext(c).(*model.LikeArticleRequest)
	if req.UserID == 0 || req.ArticleID == 0 {
		setAPIResponse(c, nil, "参数有误", false)
		return
	}
	err := service.LCService.PostLikeArticle(req.UserID, req.ArticleID)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	setAPIResponse(c, nil, "操作成功", true)
}

// PostDelLikeArticle post like
func PostDelLikeArticle(c *gin.Context) {
	req := getReqFromContext(c).(*model.LikeArticleRequest)
	if req.UserID == 0 || req.ArticleID == 0 {
		setAPIResponse(c, nil, "参数有误", false)
		return
	}
	err := service.LCService.PostDelLikeArticle(req.UserID, req.ArticleID)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	setAPIResponse(c, nil, "操作成功", true)
}
