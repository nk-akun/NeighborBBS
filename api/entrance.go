package api

import (
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
		if c.Request.URL.Path == "/api/register" {
			return new(model.RegisterRequest)
		} else if c.Request.URL.Path == "/api/login" {
			return new(model.LoginRequest)
		} else if c.Request.URL.Path == "/api/topics" {
			return new(model.ArticleRequest)
		} else if c.Request.URL.Path == "/api/comments" {
			return new(model.CommentRequest)
		} else if c.Request.URL.Path == "/api/topics/like" {
			return new(model.LikeArticleRequest)
		} else if c.Request.URL.Path == "/api/topics/del_like" {
			return new(model.LikeArticleRequest)
		}
		return nil
	}))
	r.Use(middleware.ReponseHandler())

	user := r.Group("/api")
	{
		// user.Use()
		user.POST("/register", guest.RegisterByEmail)
		user.POST("/login", guest.Login)

		user.POST("/topic", guest.PostArticle)
		user.GET("/topics", guest.GetArticleList)
		user.GET("/topics/:id", guest.GetArticleByID)
		user.POST("/topics/like", guest.PostLikeArticle)
		user.POST("/topics/del_like", guest.PostDelLikeArticle)

		user.POST("/comments", guest.PostComment)
		user.GET("/comments/:topic_id", guest.GetComments)
		// user.DELETE()
		user.GET("/test", guest.TestForUser)
	}

	r.Run()

}

func (rw *recoverWriter) Write(p []byte) (int, error) {
	logs.Logger.Error(string(p))
	return gin.DefaultErrorWriter.Write(p)
}
