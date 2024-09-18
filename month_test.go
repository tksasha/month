package month_test

import (
	"testing"
	"time"

	"github.com/tksasha/month"
	"gotest.tools/v3/assert"
)

func TestBegin(t *testing.T) {
	sbj := month.New(2024, 2).Begin()

	exp := time.Date(2024, time.Month(2), 1, 0, 0, 0, 0, time.UTC)

	assert.Equal(t, sbj, exp)
}

func TestEnd(t *testing.T) {
	sbj := month.New(2024, 2).End()

	exp := time.Date(2024, time.Month(2), 29, 23, 59, 59, 999999999, time.UTC)

	assert.Equal(t, sbj, exp)
}
