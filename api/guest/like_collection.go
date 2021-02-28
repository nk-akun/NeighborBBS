package guest

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/model"
	"github.com/nk-akun/NeighborBBS/service"
)

// PostLikeArticle post like
func PostLikeArticle(c *gin.Context) {
	user := service.UserService.GetCurrentUser(c)
	if user == nil {
		setAPIResponse(c, nil, "当前未登录！", false)
		return
	}
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
	user := service.UserService.GetCurrentUser(c)
	if user == nil {
		setAPIResponse(c, nil, "当前未登录！", false)
		return
	}
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

// PostFavoriteArticle post favorite
func PostFavoriteArticle(c *gin.Context) {
	user := service.UserService.GetCurrentUser(c)
	if user == nil {
		setAPIResponse(c, nil, "当前未登录！", false)
		return
	}
	req := getReqFromContext(c).(*model.FavoriteArticleRequest)
	if req.UserID == 0 || req.ArticleID == 0 {
		setAPIResponse(c, nil, "参数有误", false)
		return
	}
	err := service.LCService.PostFavoriteArticle(req.UserID, req.ArticleID)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	setAPIResponse(c, nil, "操作成功", true)
}

// PostDelFavoriteArticle post like
func PostDelFavoriteArticle(c *gin.Context) {
	user := service.UserService.GetCurrentUser(c)
	if user == nil {
		setAPIResponse(c, nil, "当前未登录！", false)
		return
	}
	req := getReqFromContext(c).(*model.FavoriteArticleRequest)
	if req.UserID == 0 || req.ArticleID == 0 {
		setAPIResponse(c, nil, "参数有误", false)
		return
	}
	err := service.LCService.PostDelFavoriteArticle(req.UserID, req.ArticleID)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	setAPIResponse(c, nil, "操作成功", true)
}

// GetUserFavorite ...
func GetUserFavorite(c *gin.Context) {
	user := service.UserService.GetCurrentUser(c)
	if user == nil {
		setAPIResponse(c, nil, "当前未登录！", false)
		return
	}

	var err error

	limit := c.DefaultQuery("limit", "5")
	sortby := c.DefaultQuery("sortby", "update_time")
	order := c.DefaultQuery("order", "desc")
	cursor := c.DefaultQuery("cursor", "2559090472000")
	limitNum, err1 := strconv.Atoi(limit)
	cursorTime, err2 := strconv.ParseInt(cursor, 10, 64)
	if err1 != nil || err2 != nil || limitNum <= 0 || order != "desc" && order != "asc" {
		err = errors.New("参数有误")
	}
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	resp, err := service.LCService.GetFavoriteArticles(user, limitNum, cursorTime, sortby, order)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	setAPIResponse(c, resp, "查询成功", true)
}
