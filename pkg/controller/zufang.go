package controller

import (
	"github.com/accelerise/monitor-api/pkg/model"
)

func QueryTopRiseZufangRecords() []model.ZufangFloatRecord {
	return model.QueryTopRiseZufang()
}

func QueryTopDecreaseZufangRecords() []model.ZufangFloatRecord {
	return model.QueryTopDecreaseZufang()
}