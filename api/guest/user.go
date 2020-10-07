package guest

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nk-akun/NeighborBBS/model"
)

// RegisterByEmail ...
func RegisterByEmail(c *gin.Context) {
	if req, exist := c.Get(model.CTXAPIReq); exist {
		fmt.Println(req)
	} else {
		fmt.Println("the request don't exist!")
	}
}

// TestForUser is the test api for user
func TestForUser(c *gin.Context) {
	fmt.Println(c.Request.URL.Query())
	params := c.Request.URL.Query()
	fmt.Println(params.Get("x"))
}
