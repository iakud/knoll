package kdsjson

import (
	"testing"
	"time"
)

type EnumType int32

func TestJson(t *testing.T) {
	var e Encoder
	ValueWriter[int32]()(&e, 3)
	ValueWriter[time.Time]()(&e, time.Now())
	var testEnum EnumType = 1
	ValueWriter[EnumType]()(&e, testEnum)
	t.Log(e.String())
}
