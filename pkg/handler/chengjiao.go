package handler

import (
	"github.com/accelerise/monitor-api/pkg/common/constant"
	"github.com/accelerise/monitor-api/pkg/common/util"
	"github.com/accelerise/monitor-api/pkg/controller"
	"github.com/gin-gonic/gin"
	"github.com/accelerise/monitor-api/pkg/model"
)

func GetRecentChengjiaos(ctx *gin.Context) {
	until := ctx.DefaultQuery("until", util.DefaultUntil)

	chengjiaos := controller.QueryChengjiao(until)

	ctx.JSON(constant.Success, gin.H{"data": chengjiaos, "err": nil})
}

func GetChengjiaosAverageGraph(ctx *gin.Context) {
	accuracy := ctx.DefaultQuery("accuracy", util.Month)
	from := ctx.DefaultQuery("from", util.DefaultFrom)
	until := ctx.DefaultQuery("until", util.DefaultUntil)
	xiaoqu := ctx.DefaultQuery("xiaoqu", "")

	totalPriceSumPoints, totalPriceAvgPoints, unitPriceAvgPoints := controller.QueryChengjiaoAverageGraph(from, until, accuracy, xiaoqu)
	result := map[string][]model.Point{"totalPriceSumPoints": totalPriceSumPoints, "totalPriceAvgPoints": totalPriceAvgPoints, "unitPriceAvgPoints": unitPriceAvgPoints}
	ctx.JSON(constant.Success, gin.H{"data": result, "err": nil})
}

func GetXiaoqus(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "")
	xiaoqus := controller.QueryXiaoqus(name)
	ctx.JSON(constant.Success, gin.H{"data": xiaoqus, "err": nil})
}
