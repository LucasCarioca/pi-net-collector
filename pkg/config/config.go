package config

import (
	"fmt"
	"github.com/LucasCarioca/pi-net-collector/pkg/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

var config *viper.Viper
var db *gorm.DB

func Init(env string) {
	config = viper.New()

	config.SetConfigName(fmt.Sprintf("config.%s", env))
	config.SetConfigType("yaml")
	config.AddConfigPath(".")
	config.SetEnvPrefix("secret")

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	for _, k := range config.AllKeys() {
		value := config.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			config.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value,"${"), "}")))
		}
	}

	host := config.GetString("datasource.host")
	user := config.GetString("datasource.user")
	password := config.GetString("datasource.password")
	dbname := config.GetString("datasource.dbname")
	port := config.GetString("datasource.port")
	sslmode := config.GetString("datasource.sslmode")
	timeZone := config.GetString("datasource.timeZone")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timeZone)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	models.Init(db)
}

func GetConfig() *viper.Viper {
	return config
}

func GetDataSource() *gorm.DB {
	return db
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(res) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}