package config

import (
	"fmt"
	"log"

	"github.com/kushanD95/traval-diary/package/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppConfig struct {
	pgHost     string
	pgUserName string
	pgPwd      string
	pgPort     string
	pgDB       string
	pgSSLMode  string

	AppPort string
	logger  *zap.Logger
}

var AppConfigutarion *AppConfig

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
	config.pgHost = viper.GetString("PGHOST")
	config.pgUserName = viper.GetString("PGUSRNAME")
	config.pgPwd = viper.GetString("PGPWD")
	config.pgPort = viper.GetString("PGPORT")
	config.pgDB = viper.GetString("PGDB")
	config.pgSSLMode = viper.GetString("PGSSLMODE")

}

func (config *AppConfig) InitLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf(utils.Init_error, err.Error())
		return
	}
	config.logger = logger
}

func (config *AppConfig) SetupDB() *gorm.DB {

	connStr := fmt.Sprintf(utils.DBDsn, config.pgHost, config.pgUserName, config.pgPwd, config.pgDB, config.pgPort, config.pgSSLMode)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (config *AppConfig) GetLogger() zap.Logger {
	fmt.Println(config)
	return *config.logger
}
