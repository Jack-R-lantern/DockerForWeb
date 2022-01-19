package controller

import (
	"github.com/gin-gonic/gin"
	"golang_framework/service"
)

func GetContainerList(c *gin.Context) {
	container := service.GetContainerList()
	c.JSON(200, container)
}
