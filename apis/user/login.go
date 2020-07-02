package user

import (
	"github.com/gin-gonic/gin"
	"kcsheji/model"
	"kcsheji/utils/rsps"
)

func Login(c *gin.Context) {
	number := c.PostForm("number")
	pass := c.PostForm("password")
	user := model.Users{
		Number: number,
		Password:   pass,
	}
	if !user.LoginCheck() {
		rsps.ErrLogin(c)
		return
	}
	rsps.OKWithData(c, user)
}
