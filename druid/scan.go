package druid

import (
	"github.com/gerfg/go-druid/druid/filters/types"
)

type scanQuery struct {
	QueryType    string                  `json:"queryType"`
	DataSource   DataSource              `json:"dataSource"`
	Intervals    Intervals               `json:"intervals"`
	Limit        int64                   `json:"limit"`
	ResultFormat *string                 `json:"resultFormat,omitempty"`
	Columns      *[]string               `json:"columns,omitempty"`
	Filter       *types.AggregatorFilter `json:"filter,omitempty"`
}

func NewScanQuery(
	dataSourceType string,
	dataSourceName string,
	intervals []TimeWindow,
	limit int64,
	resultFormat *string,
	columns *[]string,
	filters types.AggregatorFilter,
) *scanQuery {
	return &scanQuery{
		QueryType:    "scan",
		DataSource:   NewDataSource(dataSourceType, dataSourceName),
		Intervals:    NewIntervals(intervals),
		Limit:        limit,
		ResultFormat: resultFormat,
		Columns:      columns,
		Filter:       &filters,
	}
}
