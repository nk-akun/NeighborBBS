package guest

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/service"
)

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
		setAPIResponse(c, user, "登录成功", true)
	}
}

// TestForUser is the test api for user
func TestForUser(c *gin.Context) {
	fmt.Println(c.Request.URL.Query())
	params := c.Request.URL.Query()
	fmt.Println(params.Get("x"))
}
