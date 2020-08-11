package druid

import "time"

type DruidNativeQuery interface {
	ToString() (string, error)
}

type Interval struct {
	begin time.Time
	end   time.Time
}

func NewInterval(begin time.Time, end time.Time) Interval {
	return Interval{
		begin: begin,
		end:   end,
	}
}
