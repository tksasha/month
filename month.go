package month

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

type Month struct {
	Number int
	Name   string
	date   time.Time
}

func New(year, month string) Month {
	date, err := time.Parse(time.DateOnly, fmt.Sprintf("%04s-%02s-01", year, month))
	if err != nil {
		date = time.Now()
	}

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
		date:   date,
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

func (m Month) Begin() driver.Value {
	return m.date.AddDate(0, 0, -m.date.Day()+1).Format(time.DateOnly)
}

func (m Month) End() driver.Value {
	return m.date.AddDate(0, 1, -m.date.Day()).Format(time.DateOnly)
}
