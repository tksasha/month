package month_test

import (
	"testing"
	"time"

	"github.com/tksasha/month"
	"gotest.tools/v3/assert"
)

func TestNew(t *testing.T) {
	t.Run("returns parsed date when input is valid", func(t *testing.T) {
		month := month.New("2024", "2")

		assert.Equal(t, month.Begin, "2024-02-01")
		assert.Equal(t, month.End, "2024-02-29")
	})

	t.Run("returns current date when input is invalid", func(t *testing.T) {
		month := month.New("abc", "2")

		today := time.Now()

		begin := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, today.Location()).Format(time.DateOnly)

		end := time.Date(
			today.Year(),
			today.Month()+1,
			1,
			0,
			0,
			0,
			0,
			today.Location(),
		).Add(-1 * time.Nanosecond).Format(time.DateOnly)

		assert.Equal(t, month.Begin, begin)
		assert.Equal(t, month.End, end)
	})
}
