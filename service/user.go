package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/model"
	"github.com/nk-akun/NeighborBBS/repository"
	"github.com/nk-akun/NeighborBBS/util"
)

type userService struct {
}

// UserService is the entrance as a convenient interface
var UserService = newUserService()

func newUserService() *userService {
	return new(userService)
}

func (s *userService) SingUp(c *gin.Context) (*model.User, error) {
	req := getReqFromContext(c).(*model.RegisterRequest)
	if !util.CheckEmail(req.Email) {
		return nil, errors.New("邮箱格式错误")
	}
	if !util.CheckPassword(req.Password) {
		return nil, errors.New("密码格式错误")
	}
	if !util.CheckUsername(req.Username) {
		return nil, errors.New("用户名格式错误")
	}

	user := &model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	if err := repository.UserRepository.Create(util.DB(), user); err != nil {
		return nil, err
	}
	return user, nil
	// util.DB().Transaction(func(tx *gorm.DB) error {

	// 	return nil
	// })
}
