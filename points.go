package timeseries

import (
	"fmt"
	"time"
)

type Date struct {
	d time.Time
}

func NewDate(y int, m time.Month, d int) Date {
	return Date{
		time.Date(y, m, d, 0, 0, 0, 0, time.UTC),
	}
}

func (d Date) Before(u Date) bool {
	return d.d.Before(u.d)
}

func (d Date) String() string {
	year, month, date := d.d.Date()
	return fmt.Sprintf("%d-%d-%d", year, month, date)
}

type ComparableTime[T any] interface {
	Before(T) bool
}

func GetEarlier[C ComparableTime[C]](a, b C) C {
	if a.Before(b) {
		return a
	}
	return b
}

func GetLater[C ComparableTime[C]](a, b C) C {
	if !a.Before(b) {
		return a
	}
	return b
}
