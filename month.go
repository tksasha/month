package month

import (
	"fmt"
	"time"
)

type Month struct {
	date time.Time
}

func New(year, month string) Month {
	date, err := time.Parse(time.DateOnly, fmt.Sprintf("%04s-%02s-01", year, month))
	if err != nil {
		date = time.Now()
	}

	return Month{
		date: date,
	}
}

func (m Month) Begin() time.Time {
	return m.date
}

func (m Month) End() time.Time {
	return time.Date(
		m.date.Year(),
		m.date.Month()+1,
		m.date.Day(),
		m.date.Hour(),
		m.date.Minute(),
		m.date.Second(),
		m.date.Nanosecond(),
		m.date.Location(),
	).Add(-1 * time.Nanosecond)
}
