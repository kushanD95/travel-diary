package model

import "time"

type UserDetails struct {
	_ struct{}
	Base
	UserID    uint64    `gorm:"type:bigint;column:USER_ID;index:idx_usr_usrid;not null"`
	User      User      `gorm:"foreignKey:UserID"`
	FName     string    `gorm:"type:varchar(50);column:F_NAME"`
	LName     string    `gorm:"type:varchar(50);column:L_NAME"`
	Country   string    `gorm:"type:varchar(25);column:COUNTRY"`
	CreatedAt time.Time `gorm:"type:timestamp;column:CREATED_AT"`
}

func (UserDetails) TableName() string {
	return "USER_DETAILS"
}
