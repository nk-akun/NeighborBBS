package middleware

import (
	"bytes"
	"io/ioutil"
	"time"

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
		if err != nil {
			c.Abort()
		} else {
			c.Set(model.CTXCacheBody, requestBody)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

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
		c.Next()
	}
}

// ReponseHandler is the middleware to fill response at the end of program execution
func ReponseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()

		logs.Logger.Info("The time cost is ", end.Sub(start).Nanoseconds()/1000000)

		var resp *model.APIResponse = new(model.APIResponse)
		if c.IsAborted() {
			resp.Code = 500
			resp.Message = "Program runtime error."
			resp.Value = nil
		} else {
			resp.Code = 1000
			if message, exist := c.Get(model.CTXAPIResponseMessage); exist {
				resp.Message = message.(string)
			}
			if value, exist := c.Get(model.CTXAPIResponseValue); exist {
				resp.Value = value
			}
			value, _ := c.Get(model.CTXAPIResponseValue)
			resp.Value = value
		}
		if resp.Code == 500 {
			c.AbortWithStatusJSON(400, resp)
		}
		c.JSON(200, resp)
	}
}
