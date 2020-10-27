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
		if strings.Contains(c.Request.URL.Path, "api/register") {
			return new(model.RegisterRequest)
		} else if strings.Contains(c.Request.URL.Path, "api/login") {
			return new(model.LoginRequest)
		} else if strings.Contains(c.Request.URL.Path, "api/articles") {
			return new(model.ArticleRequest)
		}
		return nil
	}))
	r.Use(middleware.ReponseHandler())

	user := r.Group("/api")
	{
		// user.Use()
		user.POST("/register", guest.RegisterByEmail)
		user.POST("/login", guest.Login)
		user.POST("/articles", guest.PostArticle)
		user.GET("/articles", guest.GetArticleList)
		user.GET("/articles/:id", guest.GetArticleByID)
		user.POST("/comments", guest.PostComment)
		// user.DELETE()
		user.GET("/test", guest.TestForUser)
	}

	r.Run()

}

func (rw *recoverWriter) Write(p []byte) (int, error) {
	logs.Logger.Error(string(p))
	return gin.DefaultErrorWriter.Write(p)
}
