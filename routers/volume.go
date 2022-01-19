package routers

import (
	"github.com/gin-gonic/gin"
	"golang_framework/controller"
)

func initvolume(r *gin.Engine) {
	volumeApi := r.Group("/volume")
	volumeApi.GET("list", controller.VolumeList)
}
