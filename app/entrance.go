package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/logs"
)

type recoverWriter struct{}

// AppRun ...
func AppRun() {
	r := gin.New()

	r.Use(gin.RecoveryWithWriter(&recoverWriter{}))

	user := r.Group("/api/user")
	{
		user.Use()
		user.POST("/register", api.user.LoginByEmail)
	}

	r.Run()

}

func (rw *recoverWriter) Write(p []byte) (int, error) {
	logs.Logger.Error(string(p))
	return gin.DefaultErrorWriter.Write(p)
}
