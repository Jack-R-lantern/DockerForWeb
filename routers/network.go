package routers

import (
	"github.com/gin-gonic/gin"
	"golang_framework/controller"
)

func initNetwork(r *gin.Engine) {
	networkApi := r.Group("/network")
	networkApi.GET("list", controller.NetworkList)
}
