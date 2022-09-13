package config

import (
	"fmt"
	"log"

	"github.com/kushanD95/traval-diary/package/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppConfig struct {
	PGHost     string
	PGUserName string
	PGPwd      string
	PGPort     string
	PGDB       string
	PGSSLMode  string

	AppPort string
}

func (config *AppConfig) InitConfig() {

	viper.SetDefault("PGHOST", "localhost")
	viper.SetDefault("PGUSRNAME", "postgres")
	viper.SetDefault("PGPWD", "postgres")
	viper.SetDefault("PGPORT", "5431")
	viper.SetDefault("PGDB", "travel-diary")
	viper.SetDefault("PGSSLMODE", "disable")

	viper.SetDefault("APPPORT", "9000")

	viper.AutomaticEnv()

	config.AppPort = viper.GetString("APPPORT")
	config.PGHost = viper.GetString("PGHOST")
	config.PGUserName = viper.GetString("PGUSRNAME")
	config.PGPwd = viper.GetString("PGPWD")
	config.PGPort = viper.GetString("PGPORT")
	config.PGDB = viper.GetString("PGDB")
	config.PGSSLMode = viper.GetString("PGSSLMODE")

}

func (config *AppConfig) SetupDB() *gorm.DB {

	connStr := fmt.Sprintf(utils.DBDsn, config.PGHost, config.PGUserName, config.PGPwd, config.PGDB, config.PGPort, config.PGSSLMode)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
