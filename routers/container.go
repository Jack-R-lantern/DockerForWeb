package routers

import (
	"github.com/gin-gonic/gin"
	"golang_framework/controller"
)

func initContainer(r *gin.Engine) {
	containerApi := r.Group("/container")
	containerApi.GET("list", controller.GetContainerList)
}
