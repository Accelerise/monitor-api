package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Chengjiao struct {
	Url        string `json:"url"`
	XqName     string `json:"xq_name"`
	Style      string `json:"style"`
	Area       string `json:"area"`
	SignTime   string `json:"sign_time"`
	UnitPrice  string `json:"unit_price"`
	TotalPrice string `json:"total_price"`
	LngLat     string `json:"lng_lat"`
}

func QueryRecentChengjiao() []Chengjiao {

	chengjiaoDB, err := sql.Open("sqlite3", "./lianjia-detail-cj.db")
	if err != nil {
		log.Fatal(err)
	}
	defer chengjiaoDB.Close()
	rows, err := chengjiaoDB.Query("select href, name, style, area, sign_time, unit_price, total_price, lng_lat from chengjiao Limit 0,15")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	chengjiaos := make([]Chengjiao, 0)
	for rows.Next() {
		b := &Chengjiao{}
		err = rows.Scan(&b.Url, &b.XqName, &b.Style, &b.Area, &b.SignTime, &b.UnitPrice, &b.TotalPrice, &b.LngLat)
		if err != nil {
			log.Fatal(err)
		}
		chengjiaos = append(chengjiaos, *b)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return chengjiaos
}
