package month

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Month struct {
	Number int
	Name   string
}

func All() []Month {
	months := []Month{}

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

	for n, name := range names {
		months = append(months, Month{Number: n + 1, Name: name})
	}

	return months
}

func Begin(year, month string) driver.Value {
	date := date(year, month)

	return date.AddDate(0, 0, -date.Day()+1).Format(time.DateOnly)
}

func End(year, month string) driver.Value {
	date := date(year, month)

	return date.AddDate(0, 1, -date.Day()).Format(time.DateOnly)
}

func date(year, month string) time.Time {
	date, err := time.Parse(time.DateOnly, fmt.Sprintf("%04s-%02s-01", year, month))
	if err != nil {
		date = time.Now()
	}

	return date
}
