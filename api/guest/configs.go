package guest

import (
	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/model"
)

// GetConfigs return config of server
func GetConfigs(c *gin.Context) {
	resp := model.SysConfigResponse{}
	resp.SiteTitle = "碧林社区"
	resp.SiteDescription = "你内心的宁静之处"
	resp.SiteNavs = []model.ActionLink{{
		Title: "话题",
		URL:   "/topics",
	}, {
		Title: "书籍",
		URL:   "/books",
	}}
	resp.TokenExpireDays = 2
	resp.SiteKeywords = []string{"交流", "分享"}
	setAPIResponse(c, resp, "获取配置成功", true)
}
