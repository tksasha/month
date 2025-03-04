package month

import (
	"fmt"
	"strconv"
	"time"
)

type Month struct {
	Number int
	Name   string
	Begin  time.Time
	End    time.Time
}

func New(year, month string) Month {
	date, err := time.Parse(time.DateOnly, fmt.Sprintf("%04s-%02s-01", year, month))
	if err != nil {
		date = time.Now()
	}

	date = date.Truncate(time.Hour)

	names := []string{
		"Січень",
		"Лютий",
		"Березень",
		"Квітень",
		"Травень",
		"Червень",
		"Липень",
		"Серпень",
		"Вересень",
		"Жовтень",
		"Листопад",
		"Грудень",
	}

	number := int(date.Month())

	return Month{
		Number: number,
		Name:   names[number-1],
		Begin:  begin(date),
		End:    end(date),
	}
}

func All() []Month {
	months := []Month{}

	for n := range 12 {
		number := strconv.Itoa(n + 1)

		months = append(months, New("", number))
	}

	return months
}

func begin(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func end(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}
