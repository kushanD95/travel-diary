package config

import (
	"fmt"
	"log"

	"github.com/kushanD95/traval-diary/package/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppConfig struct {
	PGSSLMode  string
	PGPort     string
	PGHost     string
	PGUserName string
	PGPwd      string
	PGDB       string
}

func (config *AppConfig) SetupDB() *gorm.DB {

	config.PGHost = "localhost"
	config.PGPort = "5431"
	config.PGPwd = "postgres"
	config.PGDB = "travel-diary"
	config.PGUserName = "postgres"
	config.PGSSLMode = "disable"

	connStr := fmt.Sprintf(utils.DBDsn, config.PGHost, config.PGUserName, config.PGPwd, config.PGDB, config.PGPort, config.PGSSLMode)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
