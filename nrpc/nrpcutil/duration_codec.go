package nrpcutil

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

const maxTimeoutValue int64 = 100000000 - 1

func div(d, r time.Duration) int64 {
	if d%r > 0 {
		return int64(d/r + 1)
	}
	return int64(d / r)
}

func EncodeDuration(t time.Duration) string {
	// TODO: This is simplistic and not bandwidth efficient. Improve it.
	if t <= 0 {
		return "0n"
	}
	if d := div(t, time.Nanosecond); d <= maxTimeoutValue {
		return strconv.FormatInt(d, 10) + "n"
	}
	if d := div(t, time.Microsecond); d <= maxTimeoutValue {
		return strconv.FormatInt(d, 10) + "u"
	}
	if d := div(t, time.Millisecond); d <= maxTimeoutValue {
		return strconv.FormatInt(d, 10) + "m"
	}
	if d := div(t, time.Second); d <= maxTimeoutValue {
		return strconv.FormatInt(d, 10) + "S"
	}
	if d := div(t, time.Minute); d <= maxTimeoutValue {
		return strconv.FormatInt(d, 10) + "M"
	}
	// Note that maxTimeoutValue * time.Hour > MaxInt64.
	return strconv.FormatInt(div(t, time.Hour), 10) + "H"
}

type timeoutUnit uint8

const (
	hour        timeoutUnit = 'H'
	minute      timeoutUnit = 'M'
	second      timeoutUnit = 'S'
	millisecond timeoutUnit = 'm'
	microsecond timeoutUnit = 'u'
	nanosecond  timeoutUnit = 'n'
)

func timeoutUnitToDuration(u timeoutUnit) (d time.Duration, ok bool) {
	switch u {
	case hour:
		return time.Hour, true
	case minute:
		return time.Minute, true
	case second:
		return time.Second, true
	case millisecond:
		return time.Millisecond, true
	case microsecond:
		return time.Microsecond, true
	case nanosecond:
		return time.Nanosecond, true
	default:
	}
	return
}

func DecodeDuration(s string) (time.Duration, error) {
	size := len(s)
	if size < 2 {
		return 0, fmt.Errorf("timeout string is too short: %q", s)
	}
	if size > 9 {
		// Spec allows for 8 digits plus the unit.
		return 0, fmt.Errorf("timeout string is too long: %q", s)
	}
	unit := timeoutUnit(s[size-1])
	d, ok := timeoutUnitToDuration(unit)
	if !ok {
		return 0, fmt.Errorf("timeout unit is not recognized: %q", s)
	}
	t, err := strconv.ParseUint(s[:size-1], 10, 64)
	if err != nil {
		return 0, err
	}
	if t == 0 {
		return 0, fmt.Errorf("timeout must be positive: %q", s)
	}
	const maxHours = math.MaxInt64 / uint64(time.Hour)
	if d == time.Hour && t > maxHours {
		// This timeout would overflow math.MaxInt64; clamp it.
		return time.Duration(math.MaxInt64), nil
	}
	return d * time.Duration(t), nil
}
