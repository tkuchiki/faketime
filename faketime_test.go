package faketime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func genuineTime() time.Time {
	return time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
}

func TestNow(test *testing.T) {
	assert := assert.New(test)

	t1 := genuineTime()
	f := NewFaketime(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	f.Do()
	assert.Equal(t1, time.Now(), "not equal")
	assert.Equal(t1.Year(), time.Now().Year(), "not equal")
	assert.Equal(t1.Month(), time.Now().Month(), "not equal")
	assert.Equal(t1.Day(), time.Now().Day(), "not equal")
	assert.Equal(t1.Hour(), time.Now().Hour(), "not equal")
	assert.Equal(t1.Minute(), time.Now().Minute(), "not equal")
	assert.Equal(t1.Second(), time.Now().Second(), "not equal")
	assert.Equal(t1.Nanosecond(), time.Now().Nanosecond(), "not equal")
	assert.Equal(t1.Location(), time.Now().Location(), "not equal")
	assert.Equal(t1.YearDay(), time.Now().YearDay(), "not equal")
	assert.Equal(t1.Weekday(), time.Now().Weekday(), "not equal")
	assert.Equal(t1.UTC(), time.Now().UTC(), "not equal")
	assert.Equal(t1.Unix(), time.Now().Unix(), "not equal")
	assert.Equal(t1.UnixNano(), time.Now().UnixNano(), "not equal")
	f.Undo()

	assert.NotEqual(t1, time.Now(), "equal")
}

func TestNewFaketimeWithTime(test *testing.T) {
	assert := assert.New(test)

	t1 := genuineTime()
	f := NewFaketimeWithTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	f.Do()
	assert.Equal(t1, time.Now(), "not equal")
	assert.Equal(t1.Year(), time.Now().Year(), "not equal")
	assert.Equal(t1.Month(), time.Now().Month(), "not equal")
	assert.Equal(t1.Day(), time.Now().Day(), "not equal")
	assert.Equal(t1.Hour(), time.Now().Hour(), "not equal")
	assert.Equal(t1.Minute(), time.Now().Minute(), "not equal")
	assert.Equal(t1.Second(), time.Now().Second(), "not equal")
	assert.Equal(t1.Nanosecond(), time.Now().Nanosecond(), "not equal")
	assert.Equal(t1.Location(), time.Now().Location(), "not equal")
	assert.Equal(t1.YearDay(), time.Now().YearDay(), "not equal")
	assert.Equal(t1.Weekday(), time.Now().Weekday(), "not equal")
	assert.Equal(t1.UTC(), time.Now().UTC(), "not equal")
	assert.Equal(t1.Unix(), time.Now().Unix(), "not equal")
	assert.Equal(t1.UnixNano(), time.Now().UnixNano(), "not equal")
	f.Undo()

	assert.NotEqual(t1, time.Now(), "equal")
}
