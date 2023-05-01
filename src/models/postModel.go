package models

import (
	db "web-api/utils/database"

	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	ID    int    `gorm:"column:id; PRIMARY_KEY" json:"id"`
	Title string `gorm:"column:title;type:varchar(255);unique" json:"title"`
	Body  string `gorm:"column:body;type:varchar(255);unique" json:"body"`
	// Comments  []Comment
	// CommentID int
	// Comment   Comment `gorm:"foreignKey:CommentID"`
	UserID int
	User   Users `gorm:"foreignKey:UserID"`
}

type Comment struct {
	gorm.Model
	ID      int `gorm:"column:id; PRIMARY_KEY" json:"id"`
	UserID  int
	User    Users  `gorm:"foreignKey:UserID"`
	Content string `gorm:"column:content;type:varchar(255);unique" json:"content"`
}

func PostMigrate() {
	db.DB.Debug().AutoMigrate(&Posts{}, &Comment{})
}
