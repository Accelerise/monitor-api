package handler

import (
	"github.com/accelerise/monitor-api/pkg/common/constant"
	"github.com/accelerise/monitor-api/pkg/controller"
	"github.com/gin-gonic/gin"
)

func GetRecentChengjiaos(ctx *gin.Context) {
	chengjiaos := controller.QueryChengjiao()
	// fmt.Printf("%+v", chengjiaos)
	ctx.JSON(constant.Success, gin.H{"data": chengjiaos, "err": "error"})
}
