package user

import (
	"github.com/gin-gonic/gin"
	"kcsheji/model"
	"kcsheji/utils/rsps"
	"strconv"
)

func Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	number := c.PostForm("number")
	name := c.PostForm("name")
	class := c.PostForm("class")
	grade := c.PostForm("grade")
	phone := c.PostForm("phone")
	profession := c.PostForm("profession")
	academy := c.PostForm("academy")
	password := c.PostForm("password")
	u := model.Users{
		ID:         id,
		Number:     number,
		Name:       name,
		Class:      class,
		Grade:      grade,
		Phone:      phone,
		Profession: profession,
		Academy:    academy,
		Password:   password,
	}
	if !u.Update(u) {
		rsps.ErrUpdate(c)
		return
	}
	rsps.OK(c)
}
