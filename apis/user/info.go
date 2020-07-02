package user

import (
	"github.com/gin-gonic/gin"
	"kcsheji/model"
	"kcsheji/utils/rsps"
	"strconv"
)

func Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("user_id"))
	u := model.Users{
		ID: id,
	}
	u.GetById()
	rsps.OKWithData(c, u)
}
