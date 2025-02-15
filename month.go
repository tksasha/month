package month

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Month struct {
	Begin driver.Value
	End   driver.Value
}

func New(year, month string) Month {
	date, err := time.Parse(time.DateOnly, fmt.Sprintf("%04s-%02s-01", year, month))
	if err != nil {
		date = time.Now()
	}

	return Month{
		Begin: begin(date),
		End:   end(date),
	}
}

func begin(date time.Time) driver.Value {
	return time.Date(
		date.Year(),
		date.Month(),
		1,
		0,
		0,
		0,
		0,
		date.Location(),
	).Format(time.DateOnly)
}

func end(date time.Time) driver.Value {
	return time.Date(
		date.Year(),
		date.Month()+1,
		1,
		0,
		0,
		0,
		0,
		date.Location(),
	).Add(-1 * time.Nanosecond).Format(time.DateOnly)
}
