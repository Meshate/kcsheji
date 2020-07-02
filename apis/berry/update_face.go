package berry

import (
    "github.com/gin-gonic/gin"
    "kcsheji/model"
    "kcsheji/utils/rsps"
    "strconv"
)

func UpdateFace(c *gin.Context) {
    userId := c.PostForm("user_uuid")
    status, _ := strconv.Atoi(c.PostForm("face"))
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
    if !s.UpdateFace(status) {
        rsps.ErrUpdate(c)
        return
    }
    rsps.OK(c)
}
