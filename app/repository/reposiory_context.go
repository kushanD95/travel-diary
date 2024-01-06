package repository

import (
	"gorm.io/gorm"
)

type Context struct {
	_           struct{}
	Db          *gorm.DB
	Transaction *gorm.DB
}

func CreateRepositoryContext(db *gorm.DB, transaction *gorm.DB) Context {
	repoContext := Context{Db: db, Transaction: transaction}
	return repoContext
}
