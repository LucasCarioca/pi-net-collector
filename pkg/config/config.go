package config

import (
	"github.com/spf13/viper"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/LucasCarioca/pi-net-collector/pkg/models"
)

var config *viper.Viper
var db *gorm.DB

func Init(env string) {
	config = viper.New()

	config.SetConfigName(fmt.Sprintf("config.%s", env)) 
	config.SetConfigType("yaml")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	host := config.GetString("datasource.host")
	user := config.GetString("datasource.user")
	dbname := config.GetString("datasource.dbname")
	port := config.GetString("datasource.port")
	sslmode := config.GetString("datasource.sslmode")
	timeZone := config.GetString("datasource.timeZone")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, dbname, port, sslmode, timeZone)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	models.Init(db)
}

func GetConfig() *viper.Viper { 
	return config
}

func GetDataSource() *gorm.DB {
	return db
}