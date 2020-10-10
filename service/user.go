package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/model"
	"github.com/nk-akun/NeighborBBS/repository"
	"github.com/nk-akun/NeighborBBS/util"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
}

// UserService is the entrance as a convenient interface
var UserService = newUserService()

func newUserService() *userService {
	return new(userService)
}

func (s *userService) SignUp(c *gin.Context) (*model.User, error) {
	req := getReqFromContext(c).(*model.RegisterRequest)
	if !util.CheckEmail(req.Email) {
		return nil, errors.New("邮箱格式错误")
	} else if repository.UserRepository.GetUserByEmail(util.DB(), req.Email) != nil {
		return nil, errors.New("邮箱已被占用")
	}
	if !util.CheckUsername(req.Username) {
		return nil, errors.New("用户名格式错误")
	} else if repository.UserRepository.GetUserByUsername(util.DB(), req.Username) != nil {
		return nil, errors.New("用户名已被占用")
	}
	if !util.CheckPassword(req.Password) {
		return nil, errors.New("密码格式错误")
	}

	// encrypte password
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := &model.User{
		Username: req.Username,
		Password: string(encryptedPassword),
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

func (s *userService) Login(c *gin.Context) (*model.User, error) {
	req := getReqFromContext(c).(*model.LoginRequest)
	if req.Email == "" && req.Username == "" || req.Email != "" && req.Username != "" {
		return nil, errors.New("请使用用户名或邮箱二者之一登录")
	}
	if req.Email != "" {
		return s.loginByEmail(req.Email, req.Password)
	}
	return s.loginByUsername(req.Username, req.Password)

}

func (s *userService) loginByEmail(email string, password string) (*model.User, error) {
	user := repository.UserRepository.GetUserByEmail(util.DB(), email)
	if user == nil {
		return nil, errors.New("此邮箱不存在")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

func (s *userService) loginByUsername(username string, password string) (*model.User, error) {
	user := repository.UserRepository.GetUserByUsername(util.DB(), username)
	if user == nil {
		return nil, errors.New("此用户名不存在")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("密码错误")
	}
	return user, nil
}
