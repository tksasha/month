package month

import (
	"database/sql/driver"
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

func (m Month) Begin() (driver.Value, error) {
	return m.date.Format(time.DateOnly), nil
}

func (m Month) End() (driver.Value, error) {
	return time.Date(
		m.date.Year(),
		m.date.Month()+1,
		m.date.Day(),
		m.date.Hour(),
		m.date.Minute(),
		m.date.Second(),
		m.date.Nanosecond(),
		m.date.Location(),
	).Add(-1 * time.Nanosecond).Format(time.DateOnly), nil
}
