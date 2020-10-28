package guest

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/model"
	"github.com/nk-akun/NeighborBBS/service"
	"github.com/nk-akun/NeighborBBS/util"
)

// PostArticle is the api used to build a article
func PostArticle(c *gin.Context) {
	req := getReqFromContext(c).(*model.ArticleRequest)
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
		setAPIResponse(c, nil, err.Error())
		return
	}
	article, err := service.ArticleService.BuildArticle(req.UserID, req.Title, req.Content)
	if err != nil {
		setAPIResponse(c, nil, err.Error())
		return
	}
	setAPIResponse(c, article, "创建成功")
}

// GetArticleList is the api that returns a list of articles
// if you want to add parameter like limit, please use url like /article?limit=10
func GetArticleList(c *gin.Context) {
	limit := c.DefaultQuery("limit", "50")
	sortby := c.DefaultQuery("sortby", "id")
	order := c.DefaultQuery("order", "desc")
	// userID

	var err error

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limitNum <= 0 || order != "desc" && order != "asc" {
		err = errors.New("参数有误")
	}
	if err != nil {
		setAPIResponse(c, nil, err.Error())
	}

	resp, err := service.ArticleService.GetArticleList(limitNum, sortby, order)
	if err != nil {
		setAPIResponse(c, nil, err.Error())
	} else {
		setAPIResponse(c, resp, "查询成功")
	}
}

// GetArticleByID finds the article with ID
func GetArticleByID(c *gin.Context) {
	id := c.Param("id")

	var err error
	articleID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		setAPIResponse(c, nil, err.Error())
		return
	}

	resp, err := service.ArticleService.GetArticleByID(articleID)
	if err != nil {
		setAPIResponse(c, nil, err.Error())
		return
	}
	setAPIResponse(c, resp, "查询成功")
}
