package month_test

import (
	"testing"

	"github.com/tksasha/month"
	"gotest.tools/v3/assert"
)

func TestBegin(t *testing.T) {
	date, err := month.New("2024", "2").Begin()

	assert.NilError(t, err)
	assert.Equal(t, date, "2024-02-01")
}

func TestEnd(t *testing.T) {
	date, err := month.New("2024", "2").End()

	assert.NilError(t, err)
	assert.Equal(t, date, "2024-02-29")
}
