package time

import "time"

// ToNanoSeconds -
func ToNanoSeconds(t time.Time) int64 {
	return t.UTC().UnixNano() / int64(time.Nanosecond)
}

// ToMicroSeconds -
func ToMicroSeconds(t time.Time) int64 {
	return t.UnixNano() / int64(time.Microsecond)
}

// ToMilliSeconds -
func ToMilliSeconds(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
