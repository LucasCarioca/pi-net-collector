package server

import (
	"fmt"
	"github.com/LucasCarioca/pi-net-collector/pkg/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func routesInit(app *gin.Engine) {
	climateRouter := routes.NewClimateRouter()
	healthRouter := routes.NewHealthRouter()
	healthRouter.LinkRoutes(app)
	climateRouter.LinkRoutes(app)
}

func Init(config *viper.Viper) {
	app := gin.Default()
	app.Use(cors.Default())
	app.Use(static.Serve("/", static.LocalFile(config.GetString("server.static"), false)))
	routesInit(app)
	host := config.GetString("server.host")
	port := config.GetString("server.port")
	app.Run(fmt.Sprintf("%s:%s", host, port))
}