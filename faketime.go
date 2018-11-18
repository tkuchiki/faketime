package faketime

import (
	"time"

	"bou.ke/monkey"
)

type faketime struct {
	nowPatch *monkey.PatchGuard
	year     int
	month    time.Month
	day      int
	hour     int
	min      int
	sec      int
	nsec     int
	loc      *time.Location
	time     time.Time
}

func NewFaketime(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) *faketime {
	return &faketime{
		year:  year,
		month: month,
		day:   day,
		hour:  hour,
		min:   min,
		sec:   sec,
		nsec:  nsec,
		loc:   loc,
	}
}

func NewFaketimeWithTime(t time.Time) *faketime {
	return &faketime{
		time: t,
	}
}

func (f *faketime) Do() {
	f.nowPatch = monkey.Patch(time.Now, func() time.Time {
		if f.time.IsZero() {
			return time.Date(f.year, f.month, f.day, f.hour, f.min, f.sec, f.nsec, f.loc)
		}

		return f.time
	})
}

func (f *faketime) Undo() {
	f.nowPatch.Unpatch()
}
