package model

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	User_id      uint `gorm:"type:int(4); not null"`
	Username     string `gorm:"type:varchar(50); not null"`
	Title        string `gorm:"type:varchar(50); not null"`
	Content      string `gorm:"type:text; not null"`
	Sequence     int `gorm:"type:int(4); not null"`
	Reply_number int `gorm:"type:int(4); not null"`
	Status       int `gorm:"type:tinyint(1); not null"`
}

type PostParams struct {
Username string `form:"username" json:"username"`
Title    string `form:"title" json:"title"`
Content  string `form:"content" json:"title"`
Status   int `form:"status" json:"status"'`
}
