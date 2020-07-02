package user

import (
	"github.com/gin-gonic/gin"
	"kcsheji/model"
	"kcsheji/utils/rsps"
	"strconv"
)

func Status(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("user_id"))
	stat := model.Status{
		UserId: id,
	}
	stat.GetByUserId()
	rsps.OKWithData(c, stat)
}
