package model

import (
	"github.com/accelerise/monitor-api/pkg/common/util"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type ZufangFloatRecord struct {
	Url string `json:"url"`
	ScanAt int64 `json:"scan_at"`
	CurrentRent int64 `json:"current_price"`
	MinRent int64 `json:"min_total_price"`
	MaxRent int64 `json:"max_total_price"`
	RiseByMin int64 `json:"rise_by_min"`
	DecreaseByMax int64 `json:"decrease_by_max"`
}

func QueryTopRiseZufang() []ZufangFloatRecord {
	command := "select d.href, CAST(time_stamp as int), (CAST(rent as int)) as current_rent , " +
		"CAST(min_rent as int), CAST(max_rent as int), " +
		"(CAST(rent - min_rent as int)) as rise_by_min, (CAST(rent - max_rent as int)) as decrease_by_max " +
		"from price as a inner join " +
		"(select * from (select href, MAX(time_stamp) as latest_time from price group by href) as b inner join " +
		"(select MIN(rent) as min_rent, MAX(rent) as max_rent ,href " +
		"from price group by href) as c " +
		"on b.href = c.href) as d " +
		"on a.href = d.href and time_stamp = d.latest_time " +
		"order by rise_by_min desc limit 0,100"

	return GetZufangList(command)
}

func QueryTopDecreaseZufang() []ZufangFloatRecord {
	command := "select d.href, CAST(time_stamp as int), (CAST(rent as int)) as current_rent , " +
		"CAST(min_rent as int), CAST(max_rent as int), " +
		"(CAST(rent - min_rent as int)) as rise_by_min, (CAST(rent - max_rent as int)) as decrease_by_max " +
		"from price as a inner join " +
		"(select * from (select href, MAX(time_stamp) as latest_time from price group by href) as b inner join " +
		"(select MIN(rent) as min_rent, MAX(rent) as max_rent ,href " +
		"from price group by href) as c " +
		"on b.href = c.href) as d " +
		"on a.href = d.href and time_stamp = d.latest_time " +
		"order by decrease_by_max limit 0,100"

	return GetZufangList(command)
}

func GetZufangList(command string) []ZufangFloatRecord {
	rows := util.SqliteQuery("./lianjia-detail-zf.db", command)
	defer rows.Close()

	zufangFloatRecords := make([]ZufangFloatRecord, 0)
	for rows.Next() {
		b := ZufangFloatRecord{}
		err := rows.Scan(&b.Url, &b.ScanAt, &b.CurrentRent, &b.MinRent, &b.MaxRent, &b.RiseByMin, &b.DecreaseByMax)
		if err != nil {
			log.Fatal(err)
		}
		zufangFloatRecords = append(zufangFloatRecords, b)
	}

	return zufangFloatRecords
}