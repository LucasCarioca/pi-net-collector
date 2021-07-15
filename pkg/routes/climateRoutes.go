package routes

import (
	"github.com/LucasCarioca/pi-net-collector/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ClimateRequest struct {
	Temperature string `json:"temperature" binding:"required"`
	Humidity    string `json:"humidity" binding:"required"`
	Node        string `json:"node" binding:"required"`
	Location    string `json:"location" binding:"required"`
}

type ClimateRouter struct {
	climateService services.ClimateService
}

func NewClimateRouter() ClimateRouter {
	return ClimateRouter{
		climateService: services.NewClimateService(),
	}
}

func (r *ClimateRouter) create(context *gin.Context) {
	var climateRequest ClimateRequest
	context.BindJSON(&climateRequest)
	record := r.climateService.CreateClimateRecord(
		climateRequest.Temperature,
		climateRequest.Humidity,
		climateRequest.Node,
		climateRequest.Location)
	context.JSON(http.StatusOK, record)
}

func (r *ClimateRouter) getAll(context *gin.Context) {
	context.JSON(http.StatusOK, r.climateService.GetClimateRecords())
}

func (r *ClimateRouter) LinkRoutes(router *gin.Engine) {
	router.GET("/api/v1/climate-records", r.getAll)
	router.POST("/api/v1/climate-records", r.create)
}
