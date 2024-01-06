package model

type Base struct {
	ID uint64 `gorm:"type:bigint;column:ID;primarykey;autoincrement"`
}
