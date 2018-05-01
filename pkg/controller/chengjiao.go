package controller

import (
	"github.com/accelerise/monitor-api/pkg/model"
)

func QueryChengjiao() []model.Chengjiao {
	return model.QueryRecentChengjiao()
}
