package services

import (
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

type DBConn struct {
	Db *gorm.DB
}

func (dBConn *DBConn) SetupDB() {
	Db = dBConn.Db
}
