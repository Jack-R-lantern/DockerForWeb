package controller

import (
	"github.com/gin-gonic/gin"
	"golang_framework/service"
)

func NetworkList(c *gin.Context) {
	Networks := service.GetNetworkList()
	c.JSON(200, Networks)
}
