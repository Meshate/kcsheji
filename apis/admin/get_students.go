package admin

import (
	"github.com/gin-gonic/gin"
	"kcsheji/model"
	"kcsheji/utils/rsps"
	"strconv"
)

func GetStudents(c *gin.Context) {
	tp, _ := strconv.Atoi(c.Query("type"))
	if tp == 1 {
		var u model.Users
		var s model.Status
		users := u.SearchAll()
		stats := s.SearchAll()
		rsps.OKWithData(c, gin.H{
			"user":   users,
			"status": stats,
		})
	} else if tp == 2 {
		var u model.Users
		var s model.Status
		stats := s.SearchUndo()
		var ids []int
		for _, i := range stats {
			ids = append(ids, i.UserId)
		}
		users := u.SearchAllByIds(ids)
		rsps.OKWithData(c, gin.H{
			"user":   users,
			"status": stats,
		})
	}
}
