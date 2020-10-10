package guest

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/service"
)

// RegisterByEmail ...
func RegisterByEmail(c *gin.Context) {
	service.UserService.SingUp(c)
}

// TestForUser is the test api for user
func TestForUser(c *gin.Context) {
	fmt.Println(c.Request.URL.Query())
	params := c.Request.URL.Query()
	fmt.Println(params.Get("x"))
}
