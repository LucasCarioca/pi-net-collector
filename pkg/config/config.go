package config

import (
	"fmt"
	"github.com/LucasCarioca/pi-net-collector/pkg/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	password := config.GetString("datasource.password")
	dbname := config.GetString("datasource.dbname")
	port := config.GetString("datasource.port")
	sslmode := config.GetString("datasource.sslmode")
	timeZone := config.GetString("datasource.timeZone")
	fmt.Println(host)
	fmt.Println(user)
	fmt.Println(dbname)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timeZone)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	models.Init(db)
}

func GetConfig() *viper.Viper {
	return config
}

func GetDataSource() *gorm.DB {
	return db
}
