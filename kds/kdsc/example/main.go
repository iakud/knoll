package main

/*
#cgo LDFLAGS: -ldl -lm -Wl,-rpath,${SRCDIR}/bin ${SRCDIR}/bin/example.dylib
#include <stdlib.h>

extern int32_t apply_sync(const char* data, int32_t length);
extern char* dump(void);
*/
import "C"

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/iakud/knoll/kds/kdsc/example/kds"
)

var all = kds.NewAll(0)

func main() {
	// FIXME: init datas
	fullData, err := all.Marshal(nil)
	if err != nil {
		panic(err)
	}
	C.apply_sync(C.CString(string(fullData)), C.int32_t(len(fullData)))
	Check()
	// FIXME: change datas
	all.GetTypes().SetInt32Val(10)

	dirtyData, err := all.Marshal(nil)
	if err != nil {
		panic(err)
	}
	C.apply_sync(C.CString(string(dirtyData)), C.int32_t(len(dirtyData)))
	Check()
}

func Dump() string {
	types := all.GetTypes()
	var sb strings.Builder
	fmt.Fprintf(&sb, "Int32Val=%d\n", types.GetInt32Val())
	fmt.Fprintf(&sb, "Int64Val=%d\n", types.GetInt64Val())
	fmt.Fprintf(&sb, "Uint32Val=%d\n", types.GetUint32Val())
	fmt.Fprintf(&sb, "Uint64Val=%d\n", types.GetUint64Val())
	fmt.Fprintf(&sb, "Sint32Val=%d\n", types.GetSint32Val())
	fmt.Fprintf(&sb, "Sint64Val=%d\n", types.GetSint64Val())
	fmt.Fprintf(&sb, "Fixed32Val=%d\n", types.GetFixed32Val())
	fmt.Fprintf(&sb, "Fixed64Val=%d\n", types.GetFixed64Val())
	fmt.Fprintf(&sb, "Sfixed32Val=%d\n", types.GetSfixed32Val())
	fmt.Fprintf(&sb, "Sfixed64Val=%d\n", types.GetSfixed64Val())
	fmt.Fprintf(&sb, "FloatVal=%f\n", types.GetFloatVal())
	fmt.Fprintf(&sb, "DoubleVal=%f\n", types.GetDoubleVal())
	fmt.Fprintf(&sb, "BoolVal=%v\n", types.GetBoolVal())
	fmt.Fprintf(&sb, "StringVal=%s\n", types.GetStringVal())
	fmt.Fprintf(&sb, "BytesVal=%s\n", base64.RawStdEncoding.EncodeToString(types.GetBytesVal()))
	fmt.Fprintf(&sb, "TimestampVal=%d\n", types.GetTimestampVal().UnixNano())
	fmt.Fprintf(&sb, "DurationVal=%d\n", types.GetDurationVal().Nanoseconds())
	fmt.Fprintf(&sb, "EnumVal=%d\n", types.GetEnumVal())
	fmt.Fprintf(&sb, "ItemData=(%d,%s,%d)", types.GetItemData().GetId(), types.GetItemData().GetName(), types.GetItemData().GetCount())
	return sb.String()
}

func Check() {
	goDump := Dump()
	csharpDump := C.GoString(C.dump())
	if goDump != csharpDump {
		fmt.Printf("=== Go Dump ===\n%s\n", goDump)
		fmt.Printf("=== C# Dump ===\n%s\n", csharpDump)
		// Find first difference
		for i := 0; i < len(goDump) && i < len(csharpDump); i++ {
			if goDump[i] != csharpDump[i] {
				fmt.Printf("First diff at char %d: go='%c' (0x%x), cs='%c' (0x%x)\n", i, goDump[i], goDump[i], csharpDump[i], csharpDump[i])
				break
			}
		}
		panic("Dump mismatch!")
	}
	fmt.Println("Dump match!")
}
