package controller

import (
	"github.com/accelerise/monitor-api/pkg/model"
)

func QueryChengjiao(offset int, limit int) ([]model.Chengjiao, int) {
	return model.QueryRecentChengjiao(offset, limit)
}

func QueryChengjiaoAverageGraph(from string, until string, accuracy string, xiaoqu string) ([]model.Point, []model.Point, []model.Point) {
	return model.QueryChegnjiaoAverageGraph(from, until, accuracy, xiaoqu)
}

func QueryXiaoqus(name string) []model.Xiaoqu {
	return model.QueryXiaoqus(name)
}

func QueryDashboard(from string) model.Dashboard {
	return model.QueryDashboard(from)
}

func QueryChengjiaoMapPoint() []model.ChengjiaoMapPoint {
	return model.GetChengjiaoMapPoint()
}