package guest

import (
	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/model"
)

func setAPIResponse(c *gin.Context, value interface{}, message string, success bool) {
	if value != nil {
		c.Set(model.CTXAPIResponseValue, value)
	}
	if message != "" {
		c.Set(model.CTXAPIResponseMessage, message)
	}
	c.Set(model.CTXAPIResponseSuccess, success)
}

func getReqFromContext(c *gin.Context) interface{} {
	req, _ := c.Get(model.CTXAPIReq)
	return req
}
