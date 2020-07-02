package rsps

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
	})
}

func OKWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	})
}

func ErrBind(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 10001,
		"msg":  "request err bind",
	})
}

func ErrLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 10102,
		"msg":  "login fail",
	})
}

func ErrUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 10103,
		"msg":  "update fail",
	})
}

func ErrUpload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 10104,
		"msg":  "upload fail",
	})
}
