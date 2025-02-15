package month_test

import (
	"testing"
	"time"

	"github.com/tksasha/month"
	"gotest.tools/v3/assert"
)

func TestBegin(t *testing.T) {
	sbj := month.New("2024", "2").Begin()

	exp, err := time.Parse(time.DateOnly, "2024-02-01")
	if err != nil {
		t.Fatalf("failed to parse date: %v", err)
	}

	assert.Equal(t, sbj, exp)
}

func TestEnd(t *testing.T) {
	sbj := month.New("2024", "2").End()

	exp, err := time.Parse(time.DateTime, "2024-03-01 00:00:00")
	if err != nil {
		t.Fatalf("failed to parse date: %v", err)
	}

	exp = exp.Add(-1 * time.Nanosecond)

	assert.Equal(t, sbj, exp)
}
