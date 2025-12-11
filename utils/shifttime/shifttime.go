package shifttime

import (
	"sync"
	"sync/atomic"
	"time"
	_ "unsafe"

	"github.com/agiledragon/gomonkey/v2"
)

var locker sync.Mutex
var patchs *gomonkey.Patches

var shifted atomic.Int64

//go:linkname time_now time.now
func time_now() (sec int64, nsec int32, mono int64)

func now() time.Time {
	sec, nsec, _ := time_now()
	return time.Unix(sec, int64(nsec))
}

func Init() {
	locker.Lock()
	defer locker.Unlock()
	if patchs != nil {
		return
	}
	patchs = gomonkey.ApplyFunc(time.Now, Now)
}

func Close() {
	locker.Lock()
	defer locker.Unlock()
	if patchs == nil {
		return
	}
	patchs.Reset()
	patchs = nil
}

func Now() time.Time {
	return now().Add(time.Duration(shifted.Load()))
}

func SetTime(t time.Time) {
	shifted.Store(int64(t.Sub(now())))
}

func ShiftTime(d time.Duration) {
	shifted.Add(int64(d))
}
