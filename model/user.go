package model

import (
	"kcsheji"
	"log"
)

type Users struct {
	ID         int    `gorm:"column:id" json:"id"`
	Number     string `gorm:"column:number" json:"number"`
	Password   string `gorm:"column:password" json:"-"`
	Name       string `gorm:"column:name" json:"name"`
	Class      string `gorm:"column:class" json:"class"`
	Grade      string `gorm:"column:grade" json:"grade"`
	Phone      string `gorm:"column:phone" json:"phone"`
	Profession string `gorm:"column:profession" json:"profession"`
	Academy    string `gorm:"column:academy" json:"academy"`
	PicPath    string `gorm:"column:pic_path" json:"pic_path"`
}

func (u *Users) LoginCheck() bool {
	if err := kcsheji.Db.Where("number = ? and password = ?", u.Number, u.Password).First(u).Error; err != nil {
		return false
	}
	return true
}

func (u *Users) GetById() {
	if err := kcsheji.Db.First(u, u.ID).Error; err != nil {
		log.Println(err)
	}
}

func (u *Users) Update(user Users) bool {
	var uu Users
	if err := kcsheji.Db.First(&uu, u.ID).Updates(user).Error; err != nil {
		return false
	}
	return true
}

func (u *Users) UpdatePic(filename string) bool {
	if err := kcsheji.Db.Model(Users{}).Where("id = ?", u.ID).Update("pic_path", "/pic/"+filename).Error; err != nil {
		return false
	}
	return true
}

func (u *Users) SearchAll() (ret []Users) {
	kcsheji.Db.Find(&ret)
	return ret
}

func (u *Users) SearchAllByIds(ids []int) (ret []Users) {
	kcsheji.Db.Where("id in (?)", ids).Find(&ret)
	return ret
}

func (u *Users) GetIdByNumber() (id int) {
	var uu Users
	kcsheji.Db.Where("number = ?", u.Number).First(&uu)
	return uu.ID
}
