package models

import (
	db "web-api/utils/database"

	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	ID     int    `gorm:"column:id; PRIMARY_KEY" json:"id"`
	Title  string `gorm:"column:title;type:varchar(255)" json:"title"`
	Body   string `gorm:"column:body;type:varchar(255)" json:"body"`
	UserID int
	User   Users `gorm:"foreignKey:UserID"`
}

type Comments struct {
	gorm.Model
	ID      int    `gorm:"column:id; PRIMARY_KEY" json:"id"`
	Comment string `gorm:"column:content;type:varchar(255)" json:"content"`
	UserID  int
	User    Users `gorm:"foreignKey:UserID"`
	PostID  int
	Post    Posts `gorm:"foreignKey:PostID"`
}

func PostMigrate() {
	db.DB.Debug().AutoMigrate(&Posts{}, &Comments{})

}
