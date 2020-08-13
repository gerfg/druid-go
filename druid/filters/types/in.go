package types

type InFilter struct {
	Type      string   `json:"type"`
	Dimension string   `json:"dimension"`
	Value     []string `json:"value"`
}

func NewInFilter(dimension string, value []string) *InFilter {
	return &InFilter{
		Type:      "in",
		Dimension: dimension,
		Value:     value,
	}
}

func (filter *InFilter) ToFilter() {}
