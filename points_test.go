package timeseries

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestGetEarlier(t *testing.T) {
	t.Run("time.Time", func(t *testing.T) {
		earlier := time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC)
		later := time.Date(2022, time.August, 2, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, earlier, GetEarlier(earlier, later))
	})
	t.Run("Date", func(t *testing.T) {
		earlier := NewDate(2022, time.August, 1)
		later := NewDate(2022, time.August, 2)
		assert.Equal(t, earlier, GetEarlier(earlier, later))
	})
}