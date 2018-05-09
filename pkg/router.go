package pkg

import (
	"github.com/accelerise/monitor-api/pkg/handler"
	"github.com/gin-gonic/gin"
)

func PatchRouters(engine *gin.Engine) *gin.Engine {

	engine.GET("/chengjiaos", handler.GetRecentChengjiaos)

	engine.GET("/chengjiaos/history", handler.GetChengjiaosAverageGraph)

	engine.GET("/xiaoqus", handler.GetXiaoqus)

	engine.GET("dashboard", handler.GetDashboard)

	return engine
}
