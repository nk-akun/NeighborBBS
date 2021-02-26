package guest

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/model"
	"github.com/nk-akun/NeighborBBS/service"
)

// GetCurrentUser ...
func GetCurrentUser(c *gin.Context) {
	user := service.UserService.GetCurrentUser(c)
	setAPIResponse(c, user, "", true)
}

// GetUserInfo ...
func GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	resp, err := service.UserService.GetUserInfo(userID)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
	} else {
		setAPIResponse(c, resp, "查询成功", true)
	}
}

// RegisterByEmail ...
func RegisterByEmail(c *gin.Context) {
	user, err := service.UserService.SignUp(c)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
	} else {
		setAPIResponse(c, user, "注册成功", true)
	}
}

// Login ...
func Login(c *gin.Context) {
	user, err := service.UserService.Login(c)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
	} else {
		token := service.UserService.SetToken(user.ID)
		data := model.NewResponseValue().Set("token", token).Set("user", user)
		setAPIResponse(c, data.Value, "登录成功", true)
	}
}

// Logout ...
func Logout(c *gin.Context) {
	service.UserService.Logout(c)
	setAPIResponse(c, nil, "登出成功", true)
}

// UpdateUserProfile update nickname homePage description
func UpdateUserProfile(c *gin.Context) {
	err := service.UserService.UpdateUserProfile(c)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
	} else {
		setAPIResponse(c, nil, "操作成功", true)
	}
}

// SetUsername ...
func SetUsername(c *gin.Context) {
	err := service.UserService.SetUsername(c)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
	} else {
		setAPIResponse(c, nil, "操作成功", true)
	}
}

// SetEmail ....
func SetEmail(c *gin.Context) {
	err := service.UserService.SetEmail(c)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
	} else {
		setAPIResponse(c, nil, "操作成功", true)
	}
}

// SetPassword ...
func SetPassword(c *gin.Context) {
	err := service.UserService.SetPassword(c)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
	} else {
		setAPIResponse(c, nil, "操作成功", true)
	}
}

// UpdatePassword ...
func UpdatePassword(c *gin.Context) {
	err := service.UserService.UpdatePassword(c)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
	} else {
		setAPIResponse(c, nil, "操作成功", true)
	}
}

// TestForUser is the test api for user
func TestForUser(c *gin.Context) {
	fmt.Println(c.Request.URL.Query())
	params := c.Request.URL.Query()
	fmt.Println(params.Get("x"))
}
