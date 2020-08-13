package types

import "github.com/gerfg/go-druid/druid/filters"

type AggregatorFilter struct {
	Type   string                      `json:"type"`
	Fields []filters.DruidSingleFilter `json:"fields"`
}

func NewAggregatorFilter(fields []filters.DruidSingleFilter) AggregatorFilter {
	return AggregatorFilter{
		Type:   "and",
		Fields: fields,
	}
}

func (agg *AggregatorFilter) ToFilter() {}
