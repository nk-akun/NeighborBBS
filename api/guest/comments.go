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
	user := service.UserService.GetCurrentUser(c)
	if user == nil {
		setAPIResponse(c, nil, "当前未登录！", false)
		return
	}
	req := getReqFromContext(c).(*model.CommentRequest)
	req.Content = util.DeletePreAndSufSpace(req.Content)
	if req.UserID == 0 || req.ArticleID == 0 || req.Content == "" {
		setAPIResponse(c, nil, "参数有误", false)
		return
	}

	resp, err := service.CommentService.BuildComment(user.ID, req.ArticleID, req.ParentID, req.Content)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
	}
	setAPIResponse(c, resp, "评论成功", true)
}

// GetComments return the comments based on
func GetComments(c *gin.Context) {
	id := c.Query("article_id")
	cursor := c.DefaultQuery("cursor", "2559090472000")
	cursorTime, err1 := strconv.ParseInt(cursor, 10, 64)
	articleID, err2 := strconv.ParseInt(id, 10, 64)
	if err1 != nil || err2 != nil {
		setAPIResponse(c, nil, "参数错误", false)
		return
	}
	resp, err := service.CommentService.GetCommentList(articleID, cursorTime)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	setAPIResponse(c, resp, "查询成功", true)
}
