package controller

import (
	"github.com/accelerise/monitor-api/pkg/model"
)

func QueryTopRiseErshouRecords() []model.ErshouFloatRecord {
	return model.QueryTopRiseErshou()
}

func QueryTopDecreaseErshouRecords() []model.ErshouFloatRecord {
	return model.QueryTopDecreaseErshou()
}