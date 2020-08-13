package types

type SelectorFilter struct {
	Type      string `json:"type"`
	Dimension string `json:"dimension"`
	Value     string `json:"value"`
}

func NewselectorFilter(dimension, value string) *SelectorFilter {
	return &SelectorFilter{
		Type:      "selector",
		Dimension: dimension,
		Value:     value,
	}
}

func (filter *SelectorFilter) ToFilter() {}
