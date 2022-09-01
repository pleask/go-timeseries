package timeseries

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestCropPeriod(t *testing.T) {
	newTimePeriod := func(a, b time.Time) *TimePeriod {
		p := NewTimePeriod(a, b)
		return &p
	}
	newTime := func(i int) time.Time{
		return time.Date(2022, time.August, i+1, 0, 0, 0, 0, time.UTC)
	}
	checkCases(t, newTimePeriod, newTime)

	newDatePeriod := func(a, b Date) *DatePeriod {
		p := NewDatePeriod(a, b)
		return &p 
	}
	newDate := func(i int) Date{
		return NewDate(2022, time.August, i+1)
	}
	checkCases(t, newDatePeriod, newDate)
}

func checkCases[C ComparableTime[C], P Period[C]](t *testing.T, periodInstantiator func(C, C) P, pointInstantiator func(int) C) {
	points := make([]C, 4)
	for i := 0; i < 4; i++ {
		points[i] = pointInstantiator(i)
	}
	cases := []struct {
		periodToCrop   []int
		withinPeriod   []int
		empty bool
		expectedPeriod []int
	}{
		{
			periodToCrop: []int{0, 1},
			withinPeriod: []int{2, 3},
			empty: true,
		},
		{
			periodToCrop: []int{2, 3},
			withinPeriod: []int{0, 1},
			empty: true,
		},
		{
			periodToCrop: []int{0, 2},
			withinPeriod: []int{1, 3},
			expectedPeriod: []int{1, 2},
		},
		{
			periodToCrop: []int{1, 3},
			withinPeriod: []int{0, 2},
			expectedPeriod: []int{1, 2},
		},
		{
			periodToCrop: []int{0, 3},
			withinPeriod: []int{1, 2},
			expectedPeriod: []int{1, 2},
		},
		{
			periodToCrop: []int{1, 2},
			withinPeriod: []int{0, 3},
			expectedPeriod: []int{1, 2},
		},
		{
			periodToCrop: []int{1, 2},
			withinPeriod: []int{1, 2},
			expectedPeriod: []int{1, 2},
		},		
	}
	for i, c := range cases {
		periodToCrop := periodInstantiator(points[c.periodToCrop[0]], points[c.periodToCrop[1]])
		withinPeriod := periodInstantiator(points[c.withinPeriod[0]], points[c.withinPeriod[1]])
		CropPeriod[C](periodToCrop, withinPeriod)
		t.Run(fmt.Sprintf("%s %d", reflect.TypeOf(periodToCrop).Elem().Name(), i), func(t *testing.T) {
			if c.empty {
				assert.True(t, periodToCrop.Empty())
				return
			}
			expectedPeriod := periodInstantiator(points[c.expectedPeriod[0]], points[c.expectedPeriod[1]])
			fmt.Println(expectedPeriod, periodToCrop)
			assert.Equal(t, expectedPeriod, periodToCrop)
		})
	}
}
