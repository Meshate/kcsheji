package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"kcsheji/model"
	"kcsheji/utils/rsps"
	"os"
	"strconv"
	"time"
)

func Upload(c *gin.Context) {
	file, _, err := c.Request.FormFile("pic")
	if err != nil {
		rsps.ErrUpload(c)
		return
	}
	id, _ := strconv.Atoi(c.PostForm("user_id"))
	filename := fmt.Sprintf("%d_%d.jpg", id, time.Now().Unix())
	os.Chdir("pic")
	out, err := os.Create(filename)
	if err != nil {
		rsps.ErrUpload(c)
		return
	}
	_, _ = io.Copy(out, file)
	_ = out.Close()
	os.Chdir("..")
	u := model.Users{
		ID: id,
	}
	if !u.UpdatePic(filename) {
		rsps.ErrUpdate(c)
		return
	}
	rsps.OK(c)
}
