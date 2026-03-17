package main

/*
#cgo LDFLAGS: -ldl -lm -Wl,-rpath,${SRCDIR}/bin ${SRCDIR}/bin/example.dylib
#include <stdlib.h>

extern int32_t apply_sync(const char* data, int32_t length);
extern char* dump(void);
*/
import "C"

import (
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
	// FIXME: dump all

	return ""
}

func Check() {
	goDump := Dump()
	csharpDump := C.GoString(C.dump())
	if strings.Compare(goDump, csharpDump) != 0 {
		panic(fmt.Sprintf("go: %s, csharp: %s", goDump, csharpDump))
	}
}
