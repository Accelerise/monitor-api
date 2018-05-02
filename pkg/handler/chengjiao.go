package handler

import (
	"github.com/accelerise/monitor-api/pkg/common/constant"
	"github.com/accelerise/monitor-api/pkg/controller"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func GetRecentChengjiaos(ctx *gin.Context) {
	days, _ := time.ParseDuration("-20000h")
	// d, _ := time.ParseDuration("-24h")
	// d1 := now.Add(d)
	// fmt.Println(d1)
	until := ctx.DefaultQuery("until", strconv.FormatInt(time.Now().Add(days).Unix(), 10))
	chengjiaos := controller.QueryChengjiao(until)
	// fmt.Printf("%+v", chengjiaos)
	ctx.JSON(constant.Success, gin.H{"data": chengjiaos, "err": "error"})
}
