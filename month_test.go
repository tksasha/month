package month_test

import (
	"slices"
	"testing"

	"github.com/tksasha/month"
	"gotest.tools/v3/assert"
)

func TestAll(t *testing.T) {
	expected := []month.Month{
		{Number: 1, Name: "Січень"},
		{Number: 2, Name: "Лютий"},
		{Number: 3, Name: "Березень"},
		{Number: 4, Name: "Квітень"},
		{Number: 5, Name: "Травень"},
		{Number: 6, Name: "Червень"},
		{Number: 7, Name: "Липень"},
		{Number: 8, Name: "Серпень"},
		{Number: 9, Name: "Вересень"},
		{Number: 10, Name: "Жовтень"},
		{Number: 11, Name: "Листопад"},
		{Number: 12, Name: "Грудень"},
	}

	actual := month.All()

	assert.Assert(t, slices.Equal(actual, expected))
}

func TestBegin(t *testing.T) {
	testmap := map[string]string{
		"1":  "2024-01-01",
		"2":  "2024-02-01",
		"3":  "2024-03-01",
		"4":  "2024-04-01",
		"5":  "2024-05-01",
		"6":  "2024-06-01",
		"7":  "2024-07-01",
		"8":  "2024-08-01",
		"9":  "2024-09-01",
		"10": "2024-10-01",
		"11": "2024-11-01",
		"12": "2024-12-01",
	}

	for number, expected := range testmap {
		actual := month.Begin("2024", number)

		assert.Equal(t, actual, expected)
	}
}

func TestEnd(t *testing.T) {
	testmap := map[string]string{
		"1":  "2024-01-31",
		"2":  "2024-02-29",
		"3":  "2024-03-31",
		"4":  "2024-04-30",
		"5":  "2024-05-31",
		"6":  "2024-06-30",
		"7":  "2024-07-31",
		"8":  "2024-08-31",
		"9":  "2024-09-30",
		"10": "2024-10-31",
		"11": "2024-11-30",
		"12": "2024-12-31",
	}

	for number, expected := range testmap {
		actual := month.End("2024", number)

		assert.Equal(t, actual, expected)
	}
}
