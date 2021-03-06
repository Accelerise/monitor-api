package handler

import (
	"github.com/accelerise/monitor-api/pkg/common/constant"
	"github.com/accelerise/monitor-api/pkg/common/util"
	"github.com/accelerise/monitor-api/pkg/controller"
	"github.com/gin-gonic/gin"
	"github.com/accelerise/monitor-api/pkg/model"
	"time"
	"strconv"
)

func GetRecentChengjiaosHandler(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	chengjiaos, total := controller.QueryChengjiao(offset, limit)

	ctx.JSON(constant.Success, gin.H{"data": chengjiaos, "pagination": map[string]int{"offset": offset, "limit": limit, "total": total } })
}

func GetChengjiaosAverageGraphHandler(ctx *gin.Context) {
	accuracy := ctx.DefaultQuery("accuracy", util.Month)
	from := ctx.DefaultQuery("from", util.DefaultFrom)
	until := ctx.DefaultQuery("until", util.DefaultUntil)
	xiaoqu := ctx.DefaultQuery("xiaoqu", "")

	totalPriceSumPoints, totalPriceAvgPoints, unitPriceAvgPoints := controller.QueryChengjiaoAverageGraph(from, until, accuracy, xiaoqu)
	result := map[string][]model.Point{"totalPriceSumPoints": totalPriceSumPoints, "totalPriceAvgPoints": totalPriceAvgPoints, "unitPriceAvgPoints": unitPriceAvgPoints}
	ctx.JSON(constant.Success, gin.H{"data": result, "err": nil})
}

func GetXiaoqusHandler(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "")
	xiaoqus := controller.QueryXiaoqus(name)
	ctx.JSON(constant.Success, gin.H{"data": xiaoqus, "err": nil})
}

func GetDashboardHandler(ctx *gin.Context) {
	var AddDays, _ = time.ParseDuration("-7200h")
	var DefaultFrom = strconv.FormatInt(time.Now().Add(AddDays).Unix(), 10)
	from := ctx.DefaultQuery("from", DefaultFrom)

	dashboard := controller.QueryDashboard(from)

	ctx.JSON(constant.Success, gin.H{"data": dashboard, "err": nil})
}

func GetChengjiaoMapPointHandler(ctx *gin.Context) {
	percentl, _ := strconv.Atoi(ctx.DefaultQuery("percentl", "0"))
	percentr, _ := strconv.Atoi(ctx.DefaultQuery("percentr", "25"))

	from := ctx.DefaultQuery("from", util.DefaultFrom)
	until := ctx.DefaultQuery("until", util.DefaultUntil)
	mapPoints, minPrice := controller.QueryChengjiaoMapPoint(percentl, percentr, from, until)
	ctx.JSON(constant.Success, gin.H{"data": mapPoints, "min_price": minPrice, "err": nil})
}

func GetDistrictChengjiaoHandler(ctx *gin.Context) {
	from := ctx.DefaultQuery("from", util.DefaultFrom)
	until := ctx.DefaultQuery("until", util.DefaultUntil)
	district := ctx.DefaultQuery("district", "")
	chengjiaoStats := make([]model.DistrictStat, 0)
	if district == "" {
		chengjiaoStats = controller.QueryDistrictChengjiaoStat(from, until)
	} else {
		chengjiaoStats = controller.QueryRegionChengjiaoStat(district, from, until)
	}
	ctx.JSON(constant.Success, gin.H{"data": chengjiaoStats, "err": nil})
}

func GetRegionsByDistrictHandler(ctx *gin.Context) {
	district := ctx.Param("district")
	chengjiaoStats := controller.QueryRegionsByDistrict(district)
	ctx.JSON(constant.Success, gin.H{"data": chengjiaoStats, "err": nil})
}