package config

import (
	"github.com/kushanD95/traval-diary/app/model"
	"gorm.io/gorm"
)

func MigrateData(db *gorm.DB) {
	db.AutoMigrate(
		&model.EnvConfig{},
		&model.User{},
		&model.UserDetails{},
	)

	addedEnvDataToTable(db)
}

func addedEnvDataToTable(db *gorm.DB) {
	envs := &model.EnvConfig{EnvName: "dev", EnvURL: "dev-url"}

	db.Create(envs)
}
