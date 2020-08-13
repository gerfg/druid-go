package druid

import (
	"fmt"
	"time"
)

type DruidNativeQuery interface {
	ToDruid()
}

type Intervals struct {
	Type      string   `json:"type"`
	Intervals []string `json:"intervals"`
}

type TimeWindow struct {
	begin time.Time
	end   time.Time
}

func NewIntervals(intervals []TimeWindow) Intervals {
	stringIntervals := make([]string, 0)

	for _, window := range intervals {
		stringIntervals = append(stringIntervals, fmt.Sprintf(`%s/%s`, window.begin.Format(time.RFC3339), window.end.Format(time.RFC3339)))
	}

	return Intervals{
		Type:      "intervals",
		Intervals: stringIntervals,
	}
}

func NewTimeWindow(begin, end time.Time) TimeWindow {
	return TimeWindow{
		begin: begin,
		end:   end,
	}
}

type DataSource struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func NewDataSource(typ, name string) DataSource {
	return DataSource{
		Type: typ,
		Name: name,
	}
}
