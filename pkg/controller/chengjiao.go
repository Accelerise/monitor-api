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

func QueryChengjiaoMapPoint(percentl int, percentr int,from string, until string) ([]model.ChengjiaoMapPoint, int64) {
	return model.GetChengjiaoMapPoint(percentl, percentr, from, until)
}

func QueryDistrictChengjiaoStat(from string, until string) []model.DistrictStat {
	return model.GetDistrictChengjiaoStat(from, until)
}

func QueryRegionChengjiaoStat(district string, from string, until string) []model.DistrictStat {
	return model.GetRegionChengjiaoStat(district, from, until)
}

func QueryRegionsByDistrict(district string) []string {
	return model.GetRegionsByDistrict(district)
}