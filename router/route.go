package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	. "kcsheji"
	"kcsheji/apis/admin"
	"kcsheji/apis/berry"
	"kcsheji/apis/user"
	"net/http"
)

//路由初始化
func Init() {
	//设置正式环境
	gin.SetMode("release")

	G = gin.Default()
	G.Use(cors.Default())

	G.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	G.StaticFS("/pic", http.Dir("pic"))

	g1 := G.Group("/user")
	{
		g1.POST("/login", user.Login)
		g1.GET("/info", user.Info)
		g1.GET("/status", user.Status)
		g1.POST("/update", user.Update)
		g1.POST("/update/card", user.UpdateCard)
		g1.POST("/update/face_check", user.UpdateFace)
		g1.POST("/upload", user.Upload)
	}
	g2 := G.Group("/admin")
	{
		g2.POST("/login", admin.Login)
		g2.GET("/get_students", admin.GetStudents)
	}
	g3 := G.Group("/berry")
	{
		g3.POST("/update_card", berry.UpdateCard)
		g3.POST("/update_temp", berry.UpdateTemp)
		g3.POST("/update_face", berry.UpdateFace)
	}
}
