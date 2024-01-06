package model

import "time"

type User struct {
	_ struct{}
	Base
	UserName  string    `gorm:"type:varchar(25);column:USER_NAME"`
	Pwd       string    `gorm:"type:varchar(1000);column:PWD"`
	CreatedAt time.Time `gorm:"type:timestamp;column:CREATED_AT"`
}

func (User) TableName() string {
	return "USER"
}
