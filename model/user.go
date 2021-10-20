package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50); not null"`
	Password string `gorm:"type:char(32); not null"`
	Email    string `gorm:"type:varchar(50); not null"`
	// Birthday string `gorm:"type:date; not null"`
	Sex          int    `gorm:"type:tinyint(1); not null"`
	Address      string `gorm:"type:varchar(255); not null"`
	Post_number  int    `gorm:"type:int(4); not null"`
	Reply_number int    `gorm:"type:int(4); not null"`
	Status       int    `gorm:"type:tinyint(1); not null"`
}
