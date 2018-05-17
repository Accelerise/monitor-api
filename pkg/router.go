package pkg

import (
	"github.com/accelerise/monitor-api/pkg/handler"
	"github.com/gin-gonic/gin"
)

func PatchRouters(engine *gin.Engine) *gin.Engine {

	engine.GET("/chengjiaos", handler.GetRecentChengjiaosHandler)

	engine.GET("/chengjiaos/history", handler.GetChengjiaosAverageGraphHandler)

	engine.GET("/xiaoqus", handler.GetXiaoqusHandler)

	engine.GET("dashboard", handler.GetDashboardHandler)

	engine.GET("/ershous/top_rise", handler.GetTopRiseErshouRecordsHandler)

	engine.GET("/ershous/top_decrease", handler.GetTopDecreaseErshouRecordsHandler)

	engine.GET("/zufangs/top_rise", handler.GetTopRiseZufangRecordsHandler)

	engine.GET("/zufangs/top_decrease", handler.GetTopDecreaseZufangRecordsHandler)

	engine.GET("/chengjiaos/map", handler.GetChengjiaoMapPointHandler)

	engine.GET("/chengjiaos/district_stat", handler.GetDistrictChengjiaoHandler)

	engine.GET("/district/:district", handler.GetRegionsByDistrictHandler)

	return engine
}
