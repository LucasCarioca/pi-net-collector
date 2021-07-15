package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthRouter struct{}

func NewHealthRouter() HealthRouter {
	return HealthRouter{}
}

func (r *HealthRouter) get(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (r *HealthRouter) LinkRoutes(router *gin.Engine) {
	router.GET("/api/v1/health", r.get)
}
