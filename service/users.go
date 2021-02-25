package service

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/logs"
	"github.com/nk-akun/NeighborBBS/model"
	"github.com/nk-akun/NeighborBBS/repository"
	"github.com/nk-akun/NeighborBBS/util"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
}

// UserService is the entrance as a convenient interface
var UserService = newUserService()

func newUserService() *userService {
	return new(userService)
}

// GetCurrentUser ...
func (s *userService) GetCurrentUser(c *gin.Context) *model.User {
	token := s.GetToken(c)
	userToken, err := repository.UserTokenRepository.GetUserIDByToken(util.DB(), token)
	if err != nil {
		logs.Logger.Errorf("数据库查询token出错")
		return nil
	}
	if userToken == nil || userToken.Status == true || userToken.ExpiredAt < util.NowTimestamp() { // 不存在或者过期了
		return nil
	}
	user, err := repository.UserRepository.GetUserByUserID(util.DB(), userToken.UserID)
	if err != nil {
		logs.Logger.Errorf("数据库查询user出错")
		return nil
	}
	return user
}

func (s *userService) GetToken(c *gin.Context) string {
	token := c.GetHeader("X-User-Token")
	return token
}

func (s *userService) SignUp(c *gin.Context) (*model.User, error) {
	req := getReqFromContext(c).(*model.RegisterRequest)

	// data verification
	if !util.CheckEmail(req.Email) {
		return nil, errors.New("邮箱格式错误")
	}
	userInfo, err := repository.UserRepository.GetUserByEmail(util.DB(), req.Email)
	if err != nil {
		return nil, errors.New("查询邮箱出错")
	}
	if userInfo.ID != 0 {
		return nil, errors.New("邮箱已被占用")
	}
	if !util.CheckUsername(req.Username) {
		return nil, errors.New("用户名格式错误")
	}
	userInfo, err = repository.UserRepository.GetUserByUsername(util.DB(), req.Username)
	if err != nil {
		return nil, errors.New("查询用户名出错")
	}
	if userInfo.ID != 0 {
		return nil, errors.New("用户名已被占用")
	}
	if !util.CheckPassword(req.Password) {
		return nil, errors.New("密码格式错误")
	}

	// encrypte password
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	// get avatar url
	var avatarURL string
	if req.Email != "" {
		avatarURL = util.RandomAvatarURL(req.Email)
	} else {
		avatarURL = util.RandomAvatarURL(req.Username)
	}

	user := &model.User{
		Username:   req.Username,
		Password:   string(encryptedPassword),
		Email:      req.Email,
		Nickname:   req.Username,
		AvatarURL:  avatarURL,
		CreateTime: util.NowTimestamp(),
	}

	if err := repository.UserRepository.Create(util.DB(), user); err != nil {
		logs.Logger.Error("db error:", err)
		return nil, errors.New("数据库操作出错")
	}
	return user, nil
	// util.DB().Transaction(func(tx *gorm.DB) error {

	// 	return nil
	// })
}

func (s *userService) Login(c *gin.Context) (user *model.User, err error) {
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
	user, err := repository.UserRepository.GetUserByEmail(util.DB(), email)
	if err != nil {
		return nil, errors.New("此邮箱不存在")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

func (s *userService) loginByUsername(username string, password string) (*model.User, error) {
	user, err := repository.UserRepository.GetUserByUsername(util.DB(), username)
	if err != nil {
		return nil, errors.New("此用户名不存在")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

func (s *userService) Logout(c *gin.Context) error {
	token := s.GetToken(c)
	return repository.UserTokenRepository.UpdateStatusInvalidByToken(util.DB(), token)
}

func (s *userService) SetToken(userID int64) string {
	token := uuid.NewV4().String()
	expireTime := time.Now().Add(time.Hour * 24 * time.Duration(model.TokenExpireDays))
	userToken := &model.UserToken{
		UserID:     userID,
		Token:      token,
		ExpiredAt:  util.Timestamp(expireTime),
		CreateTime: util.NowTimestamp(),
	}
	repository.UserTokenRepository.Create(util.DB(), userToken)
	return token
}

func (s *userService) UpdateUserProfile(c *gin.Context) error {
	user := s.GetCurrentUser(c)
	if user == nil {
		return errors.New("当前未登录！")
	}
	req := getReqFromContext(c).(*model.UpdateUserProfile)
	if user.ID != req.UserID {
		return errors.New("非当前登录用户")
	}
	mp := map[string]interface{}{
		"nickname":    req.Nickname,
		"description": req.Description,
	}
	err := repository.UserRepository.UpdateMulti(util.DB(), user.ID, mp)
	if err != nil {
		logs.Logger.Errorf("数据库操作出错:%+v", err)
		return errors.New("操作失败")
	}
	return nil
}

func (s *userService) SetUsername(c *gin.Context) error {
	user := s.GetCurrentUser(c)
	if user == nil {
		return errors.New("当前未登录！")
	}
	req := getReqFromContext(c).(*model.SetUsernameRequest)
	if !util.CheckUsername(req.Username) {
		return errors.New("用户名不合法！")
	}
	err := repository.UserRepository.UpdateOne(util.DB(), user.ID, "username", req.Username)
	if err != nil {
		logs.Logger.Errorf("数据库操作出错:%+v", err)
		return errors.New("操作失败")
	}
	return nil
}

func (s *userService) SetEmail(c *gin.Context) error {
	user := s.GetCurrentUser(c)
	if user == nil {
		return errors.New("当前未登录！")
	}
	req := getReqFromContext(c).(*model.SetEmailRequest)
	if !util.CheckEmail(req.Email) {
		return errors.New("邮箱不合法！")
	}
	err := repository.UserRepository.UpdateOne(util.DB(), user.ID, "email", req.Email)
	if err != nil {
		logs.Logger.Errorf("数据库操作出错:%+v", err)
		return errors.New("操作失败")
	}
	return nil
}

func (s *userService) SetPassword(c *gin.Context) error {
	user := s.GetCurrentUser(c)
	if user == nil {
		return errors.New("当前未登录！")
	}
	req := getReqFromContext(c).(*model.SetPasswordRequest)
	if !util.CheckPassword(req.Password) {
		return errors.New("密码不合法！")
	}
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	err := repository.UserRepository.UpdateOne(util.DB(), user.ID, "password", encryptedPassword)
	if err != nil {
		logs.Logger.Errorf("数据库操作出错:%+v", err)
		return errors.New("操作失败")
	}
	return nil
}

func (s *userService) UpdatePassword(c *gin.Context) error {
	user := s.GetCurrentUser(c)
	if user == nil {
		return errors.New("当前未登录！")
	}
	req := getReqFromContext(c).(*model.UpdatePasswordRequest)
	if !util.CheckPassword(req.Password) {
		return errors.New("密码不合法！")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return errors.New("原密码输入错误")
	}
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	err := repository.UserRepository.UpdateOne(util.DB(), user.ID, "password", encryptedPassword)
	if err != nil {
		logs.Logger.Errorf("数据库操作出错:%+v", err)
	}
	return err
}

// BuildUserBriefInfo ...
func BuildUserBriefInfo(user *model.User) *model.UserBriefInfo {
	if user == nil {
		return nil
	}
	userInfo := &model.UserBriefInfo{
		ID:                    user.ID,
		Username:              user.Username,
		Nickname:              user.Nickname,
		AvatarURL:             user.AvatarURL,
		Gender:                user.Gender,
		Description:           user.Description,
		AttentionCount:        user.AttentionCount,
		FavouriteArticleCount: user.FavouriteArticleCount,
		FansCount:             user.FansCount,
		PostCount:             user.PostCount,
		CommentCount:          user.CommentCount,
		Type:                  user.Type,
		City:                  user.City,
		Province:              user.Province,
		Country:               user.Country,
	}
	return userInfo
}
