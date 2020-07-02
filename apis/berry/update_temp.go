package berry

import (
	"github.com/gin-gonic/gin"
	"kcsheji/model"
	"kcsheji/utils/rsps"
)

func UpdateTemp(c *gin.Context) {
	userId := c.PostForm("user_uuid")
	temp := c.PostForm("temperature")
	//te,_ := strconv.ParseFloat(temp,64)
	u := model.Users{
		Number: userId,
	}
	id := u.GetIdByNumber()
	if id == 0 {
		rsps.ErrUpdate(c)
	}
	s := model.Status{
		UserId: id,
	}
	if !s.UpdateTemp(temp) {
		rsps.ErrUpdate(c)
		return
	}
	rsps.OK(c)
}

