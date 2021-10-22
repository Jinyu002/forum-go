package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"type:varchar(50); not null"`
	Password     string `gorm:"type:char(32); not null"`
	Email        string `gorm:"type:varchar(50); not null"`
	Birthday     string `gorm:"type:varchar(50); not null"`
	Sex          int    `gorm:"type:tinyint(1); not null"`
	Address      string `gorm:"type:varchar(255); not null"`
	Post_number  int    `gorm:"type:int(4); not null"`
	Reply_number int    `gorm:"type:int(4); not null"`
	Status       int    `gorm:"type:tinyint(1); not null"`
}

type UserParams struct {
	Username     string `form:"username" json:"username"`
	Password     string `form:"password" json:"password"`
	Confirm      string `form:"confirm" json:"confirm"`
	Email        string `form:"email" json:"email"`
	Birthday     string `form:"birthday" json:"birthday"`
	Sex          int    `form:"sex" json:"sex"`
	Province     string `form:"province" json:"province"`
	City         string `form:"city" json:"city"`
	Area         string `form:"area" json:"area"`
	Post_number  int    `form:"post_number" json:"post_number"`
	Reply_number int    `form:"reply_number" json:"reply_number"`
	Status       int    `form:"status" json:"status"`
}

type QueryParams struct {
	Page  int `form:"page" json:"page"`
	Limit int `form:"limit" json:"limit"`
}