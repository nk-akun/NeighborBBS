package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/api/guest"
	"github.com/nk-akun/NeighborBBS/logs"
	"github.com/nk-akun/NeighborBBS/middleware"
	"github.com/nk-akun/NeighborBBS/model"
)

type recoverWriter struct{}

// AppRun ...
func AppRun() {
	r := gin.New()

	r.Use(gin.RecoveryWithWriter(&recoverWriter{}))
	r.Use(middleware.JSONRequestContextHandler(func(c *gin.Context) model.APIRequest {
		if strings.Contains(c.Request.URL.Path, "api/user/register") {
			return new(model.RegisterRequest)
		}
		return nil
	}))

	user := r.Group("/api/user")
	{
		// user.Use()
		user.POST("/register", guest.RegisterByEmail)
		user.GET("/test", guest.TestForUser)
	}

	r.Run()

}

func (rw *recoverWriter) Write(p []byte) (int, error) {
	logs.Logger.Error(string(p))
	return gin.DefaultErrorWriter.Write(p)
}
