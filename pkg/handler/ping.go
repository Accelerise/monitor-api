package handler

import (
	"github.com/accelerise/monitor-api/pkg/common/constant"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(constant.Success, gin.H{"data": 233, "err": nil})

}
