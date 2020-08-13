package types

type BoundFilter struct {
	Type        string `json:"type"`
	Dimension   string `json:"dimension"`
	Lower       string `json:"lower"`
	Upper       string `json:"upper"`
	LowerStrict bool   `json:"lowerStrict"`
	UpperStrict bool   `json:"upperStrict"`
	Ordering    string `json:"ordering"`
}

func NewBoundFilter(dimension string, lower string, upper string, lowerStrict bool, upperStrict bool, ordering string) BoundFilter {
	return BoundFilter{
		Type:   "bound",
		Dimension: dimension,
		Lower: lower,
		Upper: upper,
		LowerStrict: lowerStrict,
		UpperStrict: upperStrict,
		Ordering: ordering,
	}
}

func (agg *BoundFilter) ToFilter() {}



