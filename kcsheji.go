package kcsheji

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var G *gin.Engine
var Db *gorm.DB

//数据库连接
func init() {
	user := "zhaoyangsb"
	pass := "zhaoyangsb"
	addr := "localhost:3306"
	database := "zhaoyangsb"
	str := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, addr, database)
	fmt.Println(str)
	var err error
	Db, err = gorm.Open("mysql", str)
	if err != nil {
		//panic("database connect error")
	}
	log.Println("database connected")
}
