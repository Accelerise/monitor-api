package handler

import (
	"github.com/accelerise/monitor-api/pkg/common/constant"
	"github.com/accelerise/monitor-api/pkg/controller"
	"github.com/gin-gonic/gin"
)

func GetTopRiseErshouRecords(ctx *gin.Context) {
	ershouFloatRecords := controller.QueryTopRiseErshouRecords()

	ctx.JSON(constant.Success, gin.H{"data": ershouFloatRecords })
}

func GetTopDecreaseErshouRecords(ctx *gin.Context) {
	ershouFloatRecords := controller.QueryTopDecreaseErshouRecords()

	ctx.JSON(constant.Success, gin.H{"data": ershouFloatRecords })
}