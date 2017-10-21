package newtime

import (
	"fmt"
	"strings"
	"time"
)

type Time struct {
	Time    time.Time
	Local   *time.Location
	Is_zero bool
	Error   error
}

func NewTime() (t *Time) {
	t = new(Time)
	t.Time = time.Now().Local()
	t.Local = t.Time.Location()
	return t
}

func (t *Time) Format(format string) string {
	if t.Is_zero {
		return ParseDateFormatStringZero(format)
	}

	format = ParseDateFormat(format)
	time := t.Time.Format(format)
	return time
}

func (t *Time) FormatRaw(format string) string {
	return t.Time.Format(format)
}

func (t *Time) String() string {
	return t.Format("YYYY-mm-dd HH:ii:ss")
}

func (t *Time) SetDate(date time.Time) *Time {
	t.Time = date
	return t
}

func (t *Time) SetUnix(sec int64, nsec int64) *Time {
	t.Time = time.Unix(sec, nsec).In(t.Local)
	return t
}

// Example: t.Parse("12 января 2016", "dd mm YYYY", []string{"---", "января", "февраля", "марта", ...}, []string{"---", "январь", "февраль", "март", ...})
func (t *Time) Parse(date, format string, params ...[]string) *Time {
	date = strings.Trim(date, " \n\t\r\u00a0")
	format = ParseDateFormat(format)

	if len(params) > 0 {
		for _, months := range params {
			for n, month := range months {
				date = strings.Replace(date, month, fmt.Sprintf("%02d", n), -1)
			}
		}
	}

	return t.ParseRaw(date, format)
}

func (t *Time) ParseRaw(date, format string) *Time {
	if date == "" {
		t.Time, _ = time.Parse("2006.01.02", "0001.01.01")
		t.Is_zero = true
	} else {
		var err error
		t.Time, err = time.ParseInLocation(format, date, t.Local)
		if err != nil {
			t.Is_zero = true
		} else {
			t.Is_zero = false
		}
	}

	return t
}

func (t *Time) Add(d time.Duration) *Time {
	t.Time = t.Time.Add(d)
	return t
}

func (t *Time) LastDurationDate(params ...time.Time) time.Duration {
	var this time.Time
	if len(params) > 0 {
		this = params[0]
	} else {
		this = time.Now().Local()
	}
	return t.Time.Sub(this)
}

func (t *Time) LastDuration(date, format string) (time.Duration, error) {
	format = ParseDateFormat(format)

	this, err := time.Parse(format, date)
	if err != nil {
		return 0, err
	}
	return t.Time.Sub(this), nil
}

// see http://golang.org/src/time/format.go
func ParseDateFormat(format string) string {
	format = strings.ToLower(format)

	format = strings.Replace(format, "_long_month_", "January", -1)
	format = strings.Replace(format, "_month_", "Jan", -1)

	format = strings.Replace(format, "mm", "01", -1)
	format = strings.Replace(format, "dd", "02", -1)

	format = strings.Replace(format, "m", "1", -1)
	format = strings.Replace(format, "d", "2", -1)

	format = strings.Replace(format, "yyyy", "2006", -1)

	format = strings.Replace(format, "hh", "15", -1)
	format = strings.Replace(format, "ii", "04", -1)
	format = strings.Replace(format, "ss", "05", -1)

	format = strings.Replace(format, "h", "3", -1)
	format = strings.Replace(format, "i", "4", -1)
	format = strings.Replace(format, "s", "5", -1)

	format = strings.Replace(format, "t", "T", -1)
	format = strings.Replace(format, "z", "Z", -1)
	return format
}

func ParseDateFormatStringZero(format string) string {
	format = strings.ToLower(format)

	format = strings.Replace(format, "m", "0", -1)
	format = strings.Replace(format, "d", "0", -1)
	format = strings.Replace(format, "yyyy", "0000", -1)

	format = strings.Replace(format, "h", "0", -1)
	format = strings.Replace(format, "i", "0", -1)
	format = strings.Replace(format, "s", "0", -1)

	format = strings.Replace(format, "t", "T", -1)
	format = strings.Replace(format, "z", "Z", -1)

	format = strings.Replace(format, "Z0700", "Z", -1)
	format = strings.Replace(format, "Z07:00", "Z", -1)
	format = strings.Replace(format, "Z07", "Z", -1)
	return format
}
