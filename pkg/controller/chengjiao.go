package controller

import (
	"github.com/accelerise/monitor-api/pkg/model"
)

func QueryChengjiao(until string) []model.Chengjiao {
	return model.QueryRecentChengjiao(until)
}
