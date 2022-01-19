package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
	initContainer(r)
	initImage(r)
	initNetwork(r)
	initvolume(r)
	return r
}
