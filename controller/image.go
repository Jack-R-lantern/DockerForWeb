package controller

import (
	"github.com/gin-gonic/gin"
	"golang_framework/service"
)

func ImageList(c *gin.Context) {
	ImageList := service.GetImageList()
	c.JSON(200, ImageList)
}
