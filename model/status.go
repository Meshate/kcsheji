package model

import (
	"kcsheji"
	"log"
)

type Status struct {
	ID         int    `gorm:"column:id" json:"id"`
	UserId     int    `gorm:"column:user_id" json:"user_id"`
	IsCard     int    `gorm:"column:is_card" json:"is_card"`
	IsChecked  int    `gorm:"column:is_checked" json:"is_checked"`
	Temprature string `gorm:"column:temprature" json:"temprature"`
}

func (s *Status) GetByUserId() {
	kcsheji.Db.Where("user_id = ?", s.UserId).First(s)
}

func (s *Status) UpdateCard(card int) bool {
	if err := kcsheji.Db.Model(Status{}).Where("user_id = ?", s.UserId).Update("is_card", card).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (s *Status) UpdateFace(face int) bool {
	if err := kcsheji.Db.Model(Status{}).Where("user_id = ?", s.UserId).Update("is_checked", face).Error; err != nil {
		return false
	}
	return true
}

func (s *Status) UpdateTemp(temp string) bool {
	if err := kcsheji.Db.Model(Status{}).Where("user_id = ?", s.UserId).Update("temprature", temp).Error; err != nil {
		return false
	}
	return true
}

func (s *Status) SearchUndo() (ret []Status) {
	kcsheji.Db.Where("is_checked = 0 or is_card = 0").Find(&ret)
	return ret
}

func (s *Status) SearchAll() (ret []Status) {
	kcsheji.Db.Find(&ret)
	return ret
}
