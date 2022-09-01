package timeseries

import (
	"time"
)

type Period[T any] interface {
	Start() T
	// Should check that the start is not after the end.
	SetStart(t T)
	End() T
	// Should check that the end is not before the start.
	SetEnd(t T)
	Empty() bool
}

type TimePeriod struct {
	start time.Time
	end   time.Time
}

func NewTimePeriod(start, end time.Time) TimePeriod {
	if !start.Before(end) {
		end = start
	}
	return TimePeriod{start, end}
}

func (tp TimePeriod) Start() time.Time {
	return tp.start
}

func (tp *TimePeriod) SetStart(t time.Time) {
	if !t.Before(tp.End()) {
		t = tp.End()
	}
	tp.start = t
}

func (tp TimePeriod) End() time.Time {
	return tp.end
}

func (tp *TimePeriod) SetEnd(t time.Time) {
	if !tp.Start().Before(tp.End()) {
		t = tp.Start()
	}
	tp.end = t
}

func (tp TimePeriod) Empty() bool {
	return tp.Start() == tp.End()
}

type DatePeriod struct {
	start Date
	end   Date
}

func NewDatePeriod(start, end Date) DatePeriod {
	if !start.Before(end) {
		end = start
	}
	return DatePeriod{start, end}
}

func (dp DatePeriod) Start() Date {
	return dp.start
}

func (dp *DatePeriod) SetStart(d Date) {
	dp.start = d
}

func (dp DatePeriod) End() Date {
	return dp.end
}

func (dp *DatePeriod) SetEnd(d Date) {
	dp.end = d
}

func (dp DatePeriod) Empty() bool {
	return dp.Start() == dp.End()
}

// CropPeriod crops a period to within a period.
func CropPeriod[C ComparableTime[C], P Period[C]](period, within P) {
	start := GetLater(period.Start(), within.Start())
	end := GetEarlier(period.End(), within.End())
	if !start.Before(end) {
		end = start
	}
	period.SetStart(start)
	period.SetEnd(end)
}
