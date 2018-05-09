package controller

import (
	"github.com/accelerise/monitor-api/pkg/model"
)

func QueryChengjiao(until string) []model.Chengjiao {
	return model.QueryRecentChengjiao(until)
}

func QueryChengjiaoAverageGraph(from string, until string, accuracy string, xiaoqu string) ([]model.Point, []model.Point, []model.Point) {
	return model.QueryChegnjiaoAverageGraph(from, until, accuracy, xiaoqu)
}

func QueryXiaoqus(name string) []model.Xiaoqu {
	return model.QueryXiaoqus(name)
}