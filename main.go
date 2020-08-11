package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gerfg/go-druid/druid"
	"github.com/gerfg/go-druid/druid/filters"
)

func main() {
	begin, _ := time.Parse("2006-01-02T15:04:05-0700", "-146136543-09-08T08:23:32.096Z")
	end, _ := time.Parse("2006-01-02T15:04:05-0700", "146140482-04-24T15:36:27.903Z")
	intervals := make([]druid.Interval, 0)
	for i := 0; i < 3; i++ {
		intervals = append(intervals, druid.NewInterval(begin, end))
	}

	resultFormat := "list"
	var limit int64 = 50
	var columns []string = []string{"channel", "cityName"}

	agg := filters.NewFilterAggregator(make([]filters.DruidSingleFilter, 2))
	agg.Fields[0] = &filters.SelectorFilter{Dimension: "dim1", Value: "val1"}
	agg.Fields[1] = &filters.SelectorFilter{Dimension: "dim2", Value: "val2"}

	test := druid.NewScanQuery(
		"table",
		"wikipedia",
		intervals,
		resultFormat,
		columns,
		limit,
		agg,
	)

	if err := test.CreateNativeQuery(); err != nil {
		log.Println(err)
	}

	fmt.Println(test.JSON() + "\n\n")

	json := []byte(test.JSON())
	resp, _ := druid.Query(&json)

	fmt.Println(string(*resp))
}
