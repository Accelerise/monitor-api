package main

import (
	"github.com/accelerise/monitor-api/pkg"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.DebugMode)

	engine := pkg.PatchRouters(gin.Default())

	engine.Run()
}
