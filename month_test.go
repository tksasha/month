package month_test

import (
	"testing"

	"github.com/tksasha/month"
	"gotest.tools/v3/assert"
)

func TestNew(t *testing.T) {
	month := month.New("2024", "2")

	assert.Equal(t, month.Number, 2)
	assert.Equal(t, month.Name, "Лютий")
	assert.Equal(t, month.Begin, "2024-02-01")
	assert.Equal(t, month.End, "2024-02-29")
}

func TestAll(t *testing.T) {
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

	for idx, month := range month.All() {
		assert.Equal(t, month.Number, idx+1)
		assert.Equal(t, month.Name, names[idx])
	}
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
		month := month.New("2024", number)

		actual := month.Begin

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
		month := month.New("2024", number)

		actual := month.End

		assert.Equal(t, actual, expected)
	}
}

func TestEquals(t *testing.T) {
	left := month.New("", "")

	right := month.New("", "")

	assert.Equal(t, left, right)
}
