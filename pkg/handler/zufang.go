package handler

import (
	"github.com/accelerise/monitor-api/pkg/common/constant"
	"github.com/accelerise/monitor-api/pkg/controller"
	"github.com/gin-gonic/gin"
)

func GetTopRiseZufangRecordsHandler(ctx *gin.Context) {
	zufangFloatRecords := controller.QueryTopRiseZufangRecords()

	ctx.JSON(constant.Success, gin.H{"data": zufangFloatRecords })
}

func GetTopDecreaseZufangRecordsHandler(ctx *gin.Context) {
	zufangFloatRecords := controller.QueryTopDecreaseZufangRecords()

	ctx.JSON(constant.Success, gin.H{"data": zufangFloatRecords })
}