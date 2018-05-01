package pkg

import (
	"github.com/accelerise/monitor-api/pkg/handler"
	"github.com/gin-gonic/gin"
)

func PatchRouters(engine *gin.Engine) *gin.Engine {

	engine.GET("/ping", handler.Ping)

	engine.GET("/chengjiaos", handler.GetRecentChengjiaos)

	return engine
}
