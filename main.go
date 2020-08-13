package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gerfg/go-druid/druid"
	"github.com/gerfg/go-druid/druid/filters"
	"github.com/gerfg/go-druid/druid/filters/types"
)

func main() {
	begin, _ := time.Parse("2006-01-02T15:04:05-0700", "-146136543-09-08T08:23:32.096Z")
	end, _ := time.Parse("2006-01-02T15:04:05-0700", "146140482-04-24T15:36:27.903Z")
	intervals := make([]druid.TimeWindow, 0)
	for i := 0; i < 3; i++ {
		intervals = append(intervals, druid.NewTimeWindow(begin, end))
	}

	// resultFormat := "list"
	var limit int64 = 50
	// var columns []string = []string{"channel", "cityName"}

	agg := types.NewAggregatorFilter(make([]filters.DruidSingleFilter, 5))
	agg.Fields[0] = types.NewselectorFilter("dimension", "val1")
	agg.Fields[1] = types.NewselectorFilter("dimension", "val2")
	agg.Fields[2] = types.NewInFilter("dimension", []string{"val1", "val2", "val3"})
	agg.Fields[3] = types.NewLikeFilter("dimension", "D%")
	agg.Fields[4] = types.NewBoundFilter("dimension")

	testScan := druid.NewScanQuery(
		"table",
		"wikipedia",
		intervals,
		limit,
		nil,
		nil,
		agg,
	)

	byteArray, err := json.Marshal(testScan)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byteArray))

	// if err := test.CreateNativeQuery(); err != nil {
	// 	log.Println(err)
	// }

	// fmt.Println(test.JSON() + "\n\n")

	// json := []byte(test.JSON())
	// resp, _ := druid.Query(&json)

	// fmt.Println(string(*resp))
}
