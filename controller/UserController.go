package conroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"vo"
)

type UserController struct {
}

func (userController *UserController) GetInfo(c *gin.Context) (v *vo.GetUserInfoResponsetVo) {
	body := c.Request.FormValue("a")
	fmt.Print(body)
	u := new(vo.GetUserInfoResponsetVo)
	u.Name = "xiaowu"
	return u
}
