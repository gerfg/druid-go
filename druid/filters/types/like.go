package types

type LikeFilter struct {
	Type      string `json:"type"`
	Dimension string `json:"dimension"`
	Pattern   string `json:"value"`
}

func NewLikeFilter(dimension, pattern string) *LikeFilter {
	return &LikeFilter{
		Type:      "like",
		Dimension: dimension,
		Pattern:   pattern,
	}
}

func (filter *LikeFilter) ToFilter() {}
