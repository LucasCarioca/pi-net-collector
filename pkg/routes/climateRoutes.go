package routes

import (
	"fmt"
	"github.com/LucasCarioca/pi-net-collector/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (r *ClimateRouter) readId(context *gin.Context) (*int, string) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Not a valid id",
		})
		return nil, idStr
	}
	return &id, idStr
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

func (r *ClimateRouter) get(context *gin.Context) {
	id, _ := r.readId(context)
	if id != nil {
		record := r.climateService.GetClimateRecordById(*id)
		context.JSON(http.StatusOK, record)
	}
}

func (r *ClimateRouter) query(context *gin.Context) {
	field := context.Query("field")
	value := context.Query("value")
	last := context.Query("last")
	if field != "" && value != "" && last == "true" {
		context.JSON(http.StatusOK, r.climateService.GetLastClimateRecordBy(field, value))
	} else if field != "" && value != "" {
		context.JSON(http.StatusOK, r.climateService.GetClimateRecordsBy(field, value))
	} else {
		context.JSON(http.StatusOK, r.climateService.GetClimateRecords())
	}
}

func (r *ClimateRouter) delete(context *gin.Context) {
	id, idStr := r.readId(context)
	if id != nil {
		r.climateService.DeleteClimateRecord(*id)
		context.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("deleted record #%s", idStr),
		})
	}
}

func (r *ClimateRouter) LinkRoutes(router *gin.Engine) {
	router.GET("/api/v1/climate-records", r.query)
	router.POST("/api/v1/climate-records", r.create)
	router.GET("/api/v1/climate-records/:id", r.get)
	router.DELETE("/api/v1/climate-records/:id", r.delete)
}
