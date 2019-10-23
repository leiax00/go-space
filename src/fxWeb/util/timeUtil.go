package util

import "time"

type timeUtil struct{}

var FwTimer = &timeUtil{}

func (timer *timeUtil) CalcMillis(t time.Time) int64 {
	return t.UnixNano() / 1E6
}
