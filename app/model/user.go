package model

type User struct {
	_ struct{}
	Base
	UserName string `gorm:"type:varchar(25);column:USER_NAME"`
	Pwd      string `gorm:"type:varchar(1000);column:PWD"`
}

func (User) TableName() string {
	return "USER"
}
