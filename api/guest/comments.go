package guest

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/model"
	"github.com/nk-akun/NeighborBBS/service"
	"github.com/nk-akun/NeighborBBS/util"
)

// PostComment post comments
func PostComment(c *gin.Context) {
	req := getReqFromContext(c).(*model.CommentRequest)
	req.Content = util.DeletePreAndSufSpace(req.Content)
	if req.UserID == 0 || req.ArticleID == 0 || req.Content == "" {
		setAPIResponse(c, nil, "参数有误")
		return
	}

	resp, err := service.CommentService.BuildComment(req.UserID, req.ArticleID, req.ParentID, req.Content)
	if err != nil {
		setAPIResponse(c, nil, err.Error())
	}
	setAPIResponse(c, resp, "评论成功")
}

// GetComments return the comments based on
func GetComments(c *gin.Context) {
	id := c.Param("article_id")
	articleID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		setAPIResponse(c, nil, "参数错误")
		return
	}
	resp, err := service.CommentService.GetCommentList(articleID)
	if err != nil {
		setAPIResponse(c, nil, err.Error())
		return
	}
	setAPIResponse(c, resp, "查询成功")
}
