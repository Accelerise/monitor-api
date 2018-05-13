package model

import (
	"fmt"
	"github.com/accelerise/monitor-api/pkg/common/util"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type Chengjiao struct {
	Url        string `json:"url"`
	XqName     string `json:"xq_name"`
	Style      string `json:"style"`
	Area       string `json:"area"`
	SignTime   string `json:"sign_at"`
	UnitPrice  string `json:"unit_price"`
	TotalPrice string `json:"total_price"`
	LngLat     string `json:"lng_lat"`
}

type Xiaoqu struct {
	Name string `json:"name"`
	Regionb string `json:"district"`
	Regions string `json:"region"`
	Style string `json:"style"`
	Year string `json:"year"`
}

type Point [2]int64

type Dashboard struct {
	TotalPriceAvg int64 `json:"totalPriceAvg"`
	UnitPriceAvg int64 `json:"unitPriceAvg"`
	ChengjiaoCount int64 `json:"chengjiaoCount"`
}

func QueryDashboard(from string) Dashboard {
	command := fmt.Sprintf("select CAST(AVG(total_price) as int) as total_price_avg, " +
		"CAST(AVG(unit_price) as int) as unit_price_avg, " +
		"COUNT(*) as chengjiao_count " +
		"from chengjiao " +
		"where CAST(strftime('%%s', sign_time) as int) > %s", from)

	rows := util.SqliteQuery("./lianjia-detail-cj.db", command)
	fmt.Printf(command)
	defer rows.Close()

	b := Dashboard{}
	for rows.Next() {
		err := rows.Scan(&b.TotalPriceAvg, &b.UnitPriceAvg, &b.ChengjiaoCount)
		if err != nil {
			log.Fatal(err)
		}
	}

	return b
}

func QueryRecentChengjiao(offset int, limit int) ([]Chengjiao, int) {
	command := fmt.Sprintf("select href, name, style, area, sign_time, unit_price, total_price, lng_lat from chengjiao Order By date(sign_time) DESC Limit %d,%d", offset, limit)
	rows := util.SqliteQuery("./lianjia-detail-cj.db", command)

	defer rows.Close()

	chengjiaos := make([]Chengjiao, 0)
	for rows.Next() {
		b := &Chengjiao{}
		err := rows.Scan(&b.Url, &b.XqName, &b.Style, &b.Area, &b.SignTime, &b.UnitPrice, &b.TotalPrice, &b.LngLat)
		if err != nil {
			log.Fatal(err)
		}
		chengjiaos = append(chengjiaos, *b)
	}

	command = "select count(*) from chengjiao"
	countRows := util.SqliteQuery("./lianjia-detail-cj.db", command)
	defer countRows.Close()

	count := 0
	for countRows.Next() {
		err := countRows.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
	}

	return chengjiaos, count
}

func QueryXiaoqus(name string) []Xiaoqu {
	whereCode := ""
	if name != "" {
		whereCode = fmt.Sprintf("where name like '%%%s%%'", name)
	}
	command := fmt.Sprintf("select name, regionb, regions, style, year from xiaoqu %s Limit 0,15", whereCode)
	rows := util.SqliteQuery("./lianjia-xq.db", command)
	defer rows.Close()

	xiaoqus := make([]Xiaoqu, 0)
	for rows.Next() {
		b := &Xiaoqu{}
		err := rows.Scan(&b.Name, &b.Regionb, &b.Regions, &b.Style, &b.Year)
		if err != nil {
			log.Fatal(err)
		}
		xiaoqus = append(xiaoqus, *b)
	}

	return xiaoqus
}

func QueryChegnjiaoAverageGraph(from string, until string, accuracy util.Accuracy, xiaoqu string) ([]Point, []Point, []Point) {
	// 默认以月为间隔
	groupByCode := "group by sign_time_sub"
	strftimeCode := "strftime('%Y-%m', sign_time) as sign_time_sub"
	xiaoquCode := ""
	if accuracy == util.Quarter {
		strftimeCode = "CAST(strftime('%Y',sign_time) as int) as _year, (strftime('%m',sign_time) - 1) / 3 as quarter"
		groupByCode = "group by _year,quarter"
	}
	if accuracy == util.Year {
		strftimeCode = "strftime('%Y', sign_time) as sign_time_sub"
	}
	if accuracy == util.Day {
		strftimeCode = "strftime('%Y-%m-%d', sign_time) as sign_time_sub"
	}
	if xiaoqu != "" {
		xiaoquCode = fmt.Sprintf("and name = '%s'", xiaoqu)
	}
	command := fmt.Sprintf("select %s," +
		"CAST(SUM(total_price) as int) as total_price_sum, " +
		"CAST(AVG(total_price) as int) as total_price_avg, " +
		"CAST(AVG(unit_price) as int) as unit_price_avg " +
		"from chengjiao " +
		"where CAST(strftime('%%s', sign_time) as int) > %s " +
		"and CAST(strftime('%%s', sign_time) as int) < %s %s %s", strftimeCode, from, until, xiaoquCode,groupByCode)
	fmt.Printf(command)
	rows := util.SqliteQuery("./lianjia-detail-cj.db", command)
	defer rows.Close()

	totalPriceSumPoints := make([]Point, 0)
	totalPriceAvgPoints := make([]Point, 0)
	unitPriceAvgPoints := make([]Point, 0)
	if accuracy == util.Quarter {
		for rows.Next() {
			var year int
			var quarter int
			var totalPriceSum int64
			var totalPriceAvg int64
			var unitPriceAvg int64


			err := rows.Scan(&year, &quarter, &totalPriceSum, &totalPriceAvg, &unitPriceAvg)
			if err != nil {
				log.Fatal(err)
			}

			var groupDateTime= time.Date(year, time.Month(quarter * 3 + 1), 1, 0, 0, 0, 0, time.Local)

			groupDateUnix := groupDateTime.Unix()
			var totalPriceSumPoint= [2]int64{groupDateUnix, totalPriceSum}
			var totalPriceAvgPoint= [2]int64{groupDateUnix, totalPriceAvg}
			var unitPriceAvgPoint= [2]int64{groupDateUnix, unitPriceAvg}

			totalPriceSumPoints = append(totalPriceSumPoints, totalPriceSumPoint)
			totalPriceAvgPoints = append(totalPriceAvgPoints, totalPriceAvgPoint)
			unitPriceAvgPoints = append(unitPriceAvgPoints, unitPriceAvgPoint)
		}
		return totalPriceSumPoints, totalPriceAvgPoints, unitPriceAvgPoints
	} else {
		for rows.Next() {
			var groupDate string
			var totalPriceSum int64
			var totalPriceAvg int64
			var unitPriceAvg int64
			var groupDateTime= time.Now()

			err := rows.Scan(&groupDate, &totalPriceSum, &totalPriceAvg, &unitPriceAvg)
			if err != nil {
				log.Fatal(err)
			}
			if accuracy == util.Year {
				groupDateTime, _ = time.Parse("2006", groupDate)
			}
			if accuracy == util.Month {
				groupDateTime, _ = time.Parse("2006-01", groupDate)
			}
			if accuracy == util.Day {
				groupDateTime, _ = time.Parse("2006-01-02", groupDate)
			}
			groupDateUnix := groupDateTime.Unix()
			var totalPriceSumPoint= [2]int64{groupDateUnix, totalPriceSum}
			var totalPriceAvgPoint= [2]int64{groupDateUnix, totalPriceAvg}
			var unitPriceAvgPoint= [2]int64{groupDateUnix, unitPriceAvg}

			totalPriceSumPoints = append(totalPriceSumPoints, totalPriceSumPoint)
			totalPriceAvgPoints = append(totalPriceAvgPoints, totalPriceAvgPoint)
			unitPriceAvgPoints = append(unitPriceAvgPoints, unitPriceAvgPoint)
		}
		return totalPriceSumPoints, totalPriceAvgPoints, unitPriceAvgPoints
	}


}