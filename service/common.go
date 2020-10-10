package service

import (
	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/model"
)

func getReqFromContext(c *gin.Context) interface{} {
	req, _ := c.Get(model.CTXAPIReq)
	return req
}
