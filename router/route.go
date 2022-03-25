package router

import (
	"gin_scaffold/logger"
	"github.com/gin-gonic/gin"
)

// SetupRouter 路由
func SetupRouter() *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	return r
}
