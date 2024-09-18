package month

import (
	"time"
)

type Month struct {
	date time.Time
}

func New(year, month int) Month {
	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	return Month{date: date}
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
