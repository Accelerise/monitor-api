package model

import (
	"github.com/accelerise/monitor-api/pkg/common/util"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type ErshouFloatRecord struct {
	Url string `json:"url"`
	ScanAt int64 `json:"scan_at"`
	CurrentPrice int64 `json:"current_price"`
	MinTotalPrice int64 `json:"min_total_price"`
	MaxTotalPrice int64 `json:"max_total_price"`
	RiseByMin int64 `json:"rise_by_min"`
	DecreaseByMax int64 `json:"decrease_by_max"`
}

func QueryTopRiseErshou() []ErshouFloatRecord {
	command := "select d.href, CAST(time_stamp as int), (CAST(total_price as int)) as current_price , " +
		"CAST(min_total_price as int), CAST(max_total_price as int), " +
		"(CAST(total_price - min_total_price as int)) as rise_by_min, (CAST(total_price - max_total_price as int)) as decrease_by_max " +
		"from price as a inner join " +
		"(select * from (select href, MAX(time_stamp) as latest_time from price group by href) as b inner join " +
		"(select MIN(total_price) as min_total_price, MAX(total_price) as max_total_price ,href " +
		"from price group by href) as c " +
		"on b.href = c.href) as d " +
		"on a.href = d.href and time_stamp = d.latest_time " +
		"order by rise_by_min desc limit 0,100"

	return GetErshouList(command)
}

func QueryTopDecreaseErshou() []ErshouFloatRecord {
	command := "select d.href, CAST(time_stamp as int), (CAST(total_price as int)) as current_price , " +
		"CAST(min_total_price as int), CAST(max_total_price as int), " +
		"(CAST(total_price - min_total_price as int)) as rise_by_min, (CAST(total_price - max_total_price as int)) as decrease_by_max " +
		"from price as a inner join " +
		"(select * from (select href, MAX(time_stamp) as latest_time from price group by href) as b inner join " +
		"(select MIN(total_price) as min_total_price, MAX(total_price) as max_total_price ,href " +
		"from price group by href) as c " +
		"on b.href = c.href) as d " +
		"on a.href = d.href and time_stamp = d.latest_time " +
		"order by decrease_by_max limit 0,100"

	return GetErshouList(command)
}

func GetErshouList(command string) []ErshouFloatRecord {
	rows := util.SqliteQuery("./lianjia-detail-es.db", command)
	defer rows.Close()

	ershouFloatRecords := make([]ErshouFloatRecord, 0)
	for rows.Next() {
		b := ErshouFloatRecord{}
		err := rows.Scan(&b.Url, &b.ScanAt, &b.CurrentPrice, &b.MinTotalPrice, &b.MaxTotalPrice, &b.RiseByMin, &b.DecreaseByMax)
		if err != nil {
			log.Fatal(err)
		}
		ershouFloatRecords = append(ershouFloatRecords, b)
	}

	return ershouFloatRecords
}