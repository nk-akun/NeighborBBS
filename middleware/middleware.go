package middleware

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/logs"
	"github.com/nk-akun/NeighborBBS/model"
)

// GetAPIRequestModel returns the model used to store request info
type GetAPIRequestModel func(*gin.Context) model.APIRequest

// JSONRequestContextHandler is the middleware to preproccess request
func JSONRequestContextHandler(getAPIRequestModel GetAPIRequestModel) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody, err := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		if err != nil {
			c.Abort()
		}
		params := c.Request.URL.Query()
		c.Set(model.CTXAPICacheBody, requestBody)
		c.Set(model.CTXAPIURLParams, params)
		req := getAPIRequestModel(c)
		if req != nil {
			if err = c.BindJSON(req); err == nil {
				c.Set(model.CTXAPIReq, req)
			} else {
				logs.Logger.Error("parse json error:", err)
				c.Abort()
			}
		} else {
			logs.Logger.Error("program can not find the struct matched the request!")
			c.Abort()
		}
	}
}
