package guest

import (
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

	service.CommentService.BuildComment(req.UserID, req.ArticleID, req.ParentID, req.Content)
}
