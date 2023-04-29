package models

import (
	db "web-api/utils/database"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Unique_id int    `gorm:"column:unique_id; PRIMARY_KEY" json:"unique_id"`
	Title     string `gorm:"column:title;type:varchar(255);unique" json:"title"`
	Body      string `gorm:"column:body;type:varchar(255);unique" json:"body"`
	Comments  []Comment
	User      Users `gorm:"foreignKey:Users.Unique_id"`
}

type Comment struct {
	gorm.Model
	Unique_id int `gorm:"column:unique_id; PRIMARY_KEY" json:"unique_id"`
	User      Users
	Content   string `gorm:"column:content;type:varchar(255);unique" json:"content"`
}

func PostMigrate() {
	db.DB.Debug().AutoMigrate(&Post{}, &Comment{})
}
