package controller

import (
	"github.com/gin-gonic/gin"
	"golang_framework/service"
)

func VolumeList(c *gin.Context) {
	Volumes := service.GetVolumeList()
	c.JSON(200, Volumes)
}
