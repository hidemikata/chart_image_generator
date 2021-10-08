package model

import (
	"chart_image_generator/def"
)

func GetPriceTest() []def.BtcJpy {
	past_str := "2021-09-10 10:10:00"
	latest_str := "2021-09-20 10:50:10"
	//データをDBから取得
	rows, err := db.Query(`select * from btc_jpy_live where date between "` + past_str + `" and "` + latest_str + `" order by date limit 100;`)
	if err != nil {
		panic(err.Error())
	}

	records := make([]def.BtcJpy, 0)
	for rows.Next() {
		var record def.BtcJpy
		err = rows.Scan(
			&record.Date,
			&record.Symbol,
			&record.Open,
			&record.High,
			&record.Low,
			&record.Close,
		)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, record)
	}

	return records
}
