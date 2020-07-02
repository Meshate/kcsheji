package user

import (
	"github.com/gin-gonic/gin"
	"kcsheji/model"
	"kcsheji/utils/rsps"
	"strconv"
)

func UpdateCard(c *gin.Context) {
	userId, _ := strconv.Atoi(c.PostForm("user_id"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	s := model.Status{
		UserId: userId,
	}
	if !s.UpdateCard(status) {
		rsps.ErrUpdate(c)
		return
	}
	rsps.OK(c)
}
