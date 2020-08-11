package filters

import (
	"fmt"
	"strings"
)

type DruidSingleFilter interface {
	ToFilter() (string, error)
}

type FilterAggregator struct {
	Type   string
	Fields []DruidSingleFilter
}

func NewFilterAggregator(fields []DruidSingleFilter) FilterAggregator {
	return FilterAggregator{
		Type:   "and",
		Fields: fields,
	}
}

func (agg *FilterAggregator) ToFilter() (string, error) {
	filter := fmt.Sprintf(`, "filter": {"type": "%s", "fields": [`, agg.Type)

	for _, field := range agg.Fields {
		fieldString, err := field.ToFilter()
		if err != nil {
			return "", err
		}
		filter += fieldString
	}

	filter = strings.TrimSuffix(filter, ", ")

	filter += `]}`

	return filter, nil
}
