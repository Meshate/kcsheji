package admin

import (
	"github.com/gin-gonic/gin"
	"kcsheji/utils/rsps"
)

func Login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	if account == "admin" && password == "admin" {
		rsps.OK(c)
	} else {
		rsps.ErrLogin(c)
	}
}
