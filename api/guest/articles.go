package guest

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/logs"
	"github.com/nk-akun/NeighborBBS/model"
	"github.com/nk-akun/NeighborBBS/service"
	"github.com/nk-akun/NeighborBBS/util"
)

// PostArticle is the api used to build a article
func PostArticle(c *gin.Context) {
	user := service.UserService.GetCurrentUser(c)
	if user == nil {
		setAPIResponse(c, nil, "当前未登录！", false)
		return
	}
	req := getReqFromContext(c).(*model.ArticleRequest)
	logs.Logger.Info(req)
	var err error
	if req.UserID == 0 {
		err = errors.New("user_id is invalid")
	}
	if !util.CheckContent(req.Content) {
		err = errors.New("提交失败，内容为空")
	}
	if !util.CheckContent(req.Title) {
		err = errors.New("创建失败，标题为空")
	}
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	article, err := service.ArticleService.BuildArticle(user, req.Title, req.Content)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	setAPIResponse(c, article, "创建成功", true)
}

// GetArticleList is the api that returns a list of articles
// if you want to add parameter like limit, please use url like /article?limit=10
func GetArticleList(c *gin.Context) {
	limit := c.DefaultQuery("limit", "10")
	sortby := c.DefaultQuery("sortby", "create_time")
	order := c.DefaultQuery("order", "desc")
	cursor := c.DefaultQuery("cursor", "2559090472000")
	// userID

	var err error

	limitNum, err1 := strconv.Atoi(limit)
	cursorTime, err2 := strconv.ParseInt(cursor, 10, 64)
	if err1 != nil || err2 != nil || limitNum <= 0 || order != "desc" && order != "asc" {
		err = errors.New("参数有误")
	}
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
	}

	resp, err := service.ArticleService.GetArticleList(limitNum, cursorTime, sortby, order)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
	} else {
		setAPIResponse(c, resp, "查询成功", true)
	}
}

// GetArticleByID finds the article with ID
func GetArticleByID(c *gin.Context) {
	id := c.Param("id")

	var err error
	articleID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}

	resp, err := service.ArticleService.GetArticleByID(articleID)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	setAPIResponse(c, resp, "查询成功", true)
}
