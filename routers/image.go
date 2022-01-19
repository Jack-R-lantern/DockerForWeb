package routers

import (
	"github.com/gin-gonic/gin"
	"golang_framework/controller"
)

func initImage(r *gin.Engine) {
	imageApi := r.Group("/image")
	imageApi.GET("list", controller.ImageList)
}
