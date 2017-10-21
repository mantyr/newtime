package newtime

import (
	"time"
)

// If the name is "" or "UTC", LoadLocation returns UTC. If the name is "Local", LoadLocation returns Local.
// /usr/lib/go/lib/time/zoneinfo.zip
func (t *Time) SetLocal(name string) *Time {
	t.Local, t.Error = time.LoadLocation(name)
	return t
}

func (t *Time) SetUTC() *Time {
	t.SetLocal("UTC")
	return t
}

func (t *Time) SetPST() *Time {
	t.SetLocal("PST8PDT")
	return t
}

func (t *Time) SetMoscow() *Time {
	t.SetLocal("Europe/Moscow")
	return t
}
