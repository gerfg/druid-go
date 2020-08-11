package filters

import "fmt"

type SelectorFilter struct {
	Type      string
	Dimension string
	Value     string
}

func NewselectorFilter(dimension, value string) SelectorFilter {
	return SelectorFilter{
		Type:      "selector",
		Dimension: dimension,
		Value:     value,
	}
}

func (filter *SelectorFilter) ToFilter() (string, error) {
	return fmt.Sprintf(`{"type": "selector, "dimension": "%s", "value": "%s"}, `, filter.Dimension, filter.Value), nil
}
