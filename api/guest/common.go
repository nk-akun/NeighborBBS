package guest

import (
	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/model"
)

func setAPIResponse(c *gin.Context, value interface{}, message string) {
	if value != nil {
		c.Set(model.CTXAPIResponseValue, value)
	}
	if message != "" {
		c.Set(model.CTXAPIResponseMessage, message)
	}
}

func getReqFromContext(c *gin.Context) interface{} {
	req, _ := c.Get(model.CTXAPIReq)
	return req
}
